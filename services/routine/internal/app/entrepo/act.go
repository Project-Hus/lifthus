package entrepo

import (
	"context"
	"log"
	"routine/internal/domain/aggregates/act"
	"routine/internal/ent"
	eact "routine/internal/ent/act"
	"routine/internal/ent/actimage"
	eactv "routine/internal/ent/actversion"
	"routine/internal/ent/image"
	"routine/internal/repository"
	"time"
)

func NewEntActRepository() *EntActRepository {
	return &EntActRepository{EntRepository: NewEntRepository()}
}

type EntActRepository struct {
	*EntRepository
}

func (repo *EntActRepository) FindActByCode(ctx context.Context, code act.ActCode) (fAct *act.Act, err error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	a, err := tx.Act.Query().
		Where(eact.Code(string(code))).
		WithActVersions(
			func(q *ent.ActVersionQuery) {
				q.Order(ent.Asc((eactv.FieldVersion)))
				q.WithImages()
				q.WithActImages(
					func(q *ent.ActImageQuery) {
						q.Order(ent.Asc(actimage.FieldOrder))
						q.WithImage()
					},
				)
			},
		).
		First(ctx)
	if ent.IsNotFound(err) {
		return nil, repository.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return repo.actFromEntAct(ctx, a)
}

func (repo *EntActRepository) FindActsByName(ctx context.Context, actName string) (fActs []*act.Act, err error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	eacts, err := tx.Act.Query().Where(eact.NameContains(actName)).All(ctx)
	if err != nil {
		return nil, err
	}
	fActs = make([]*act.Act, len(eacts))
	for i, eact := range eacts {
		fActs[i], err = repo.actFromEntAct(ctx, eact)
		if err != nil {
			return nil, err
		}
	}
	return fActs, nil
}

func (repo *EntActRepository) Save(ctx context.Context, target *act.Act) (sAct *act.Act, err error) {
	_, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	prev, err := repo.FindActByCode(ctx, target.Code())
	if repository.IsNotFound(err) {
		return repo.insertNewAct(ctx, target)
	} else if err != nil {
		return nil, err
	}
	return repo.updateVersions(ctx, prev, target)
}

func (repo *EntActRepository) insertNewAct(ctx context.Context, target *act.Act) (*act.Act, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}

	eat, err := entActTypeFromActType(target.Type())
	if err != nil {
		return nil, err
	}
	eact, err := tx.Act.Create().
		SetCode(string(target.Code())).
		SetActType(eat).
		SetName(string(target.Name())).
		SetAuthor(uint64(target.Author())).
		SetCreatedAt(time.Time(target.CreatedAt())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	vs, err := repo.createVerBulk(ctx, eact, target.Versions())
	if err != nil {
		return nil, err
	}
	eact.Edges.ActVersions = vs
	return repo.actFromEntAct(ctx, eact)
}

func (repo *EntActRepository) createVerBulk(ctx context.Context, eact *ent.Act, vers act.ActVersions) ([]*ent.ActVersion, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}

	states := make([]*ent.ActVersionCreate, len(vers))
	for i, v := range vers {
		states[i] = tx.ActVersion.Create().
			SetCode(string(v.Code())).SetVersion(uint(v.Version())).
			SetText(string(v.Text())).SetCreatedAt(time.Time(v.CreatedAt())).
			SetAct(eact).SetActCode(eact.Code)
	}
	evs, err := tx.ActVersion.CreateBulk(states...).Save(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("evs", evs)
	for i, ev := range evs {
		imgs, err := repo.createImgBulk(ctx, ev, vers[i].ImageSrcs())
		if err != nil {
			return nil, err
		}
		ev.Edges.ActImages = imgs
	}
	return evs, nil
}

func (repo *EntActRepository) createImgBulk(ctx context.Context, ver *ent.ActVersion, imgs act.ActImageSrcs) ([]*ent.ActImage, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	states := make([]*ent.ActImageCreate, len(imgs))
	for i, img := range imgs {
		img, err := tx.Image.Query().Where(image.Src(img)).First(ctx)
		if err != nil {
			return nil, err
		}
		states[i] = tx.ActImage.Create().SetOrder(uint(i) + 1).SetImage(img)
	}
	return tx.ActImage.CreateBulk(states...).Save(ctx)
}

func (repo *EntActRepository) updateVersions(ctx context.Context, prev *act.Act, cur *act.Act) (*act.Act, error) {
	versToDel := []*act.ActVersion{}
	versToCreate := []*act.ActVersion{}

	prevVersions := prev.VersionsMap()
	curVersions := cur.VersionsMap()

	for pvn, pv := range prevVersions {
		if _, ok := curVersions[pvn]; !ok {
			versToDel = append(versToDel, pv)
		}
	}

	for cvn, cv := range curVersions {
		if _, ok := prevVersions[cvn]; !ok {
			versToCreate = append(versToCreate, cv)
		}
	}

	err := repo.deleteVersions(ctx, versToDel)
	if err != nil {
		return nil, err
	}
	err = repo.createVersions(ctx, cur.Code(), versToCreate)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

func (repo *EntActRepository) deleteVersions(ctx context.Context, vs []*act.ActVersion) error {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return err
	}
	for _, v := range vs {
		_, err := tx.ActVersion.Delete().Where(eactv.Code(string(v.Code()))).Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *EntActRepository) createVersions(ctx context.Context, actCode act.ActCode, vs []*act.ActVersion) error {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return err
	}
	eact, err := tx.Act.Query().Where(eact.Code(string(actCode))).First(ctx)
	if err != nil {
		return err
	}
	for _, v := range vs {
		_, err = tx.ActVersion.Create().
			SetCode(string(v.Code())).SetVersion(uint(v.Version())).
			SetText(string(v.Text())).SetCreatedAt(time.Time(v.CreatedAt())).
			SetAct(eact).SetActCode(eact.Code).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
