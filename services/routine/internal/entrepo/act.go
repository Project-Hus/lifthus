package entrepo

import (
	"context"
	"log"
	"routine/internal/domain/aggregates/act"
	"routine/internal/ent"
	eact "routine/internal/ent/act"
	eacti "routine/internal/ent/actimage"
	eactv "routine/internal/ent/actversion"
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
	finally, err := repo.BeginOrContinueTx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	a, err := repo.tx.Act.Query().
		Where(eact.Code(string(code))).
		WithActVersions(
			func(q *ent.ActVersionQuery) {
				q.Order(ent.Asc((eactv.FieldVersion)))
				q.WithActImages(
					func(q *ent.ActImageQuery) {
						q.Order(ent.Asc(eacti.FieldOrder))
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

func (repo *EntActRepository) Save(ctx context.Context, target *act.Act) (sAct *act.Act, err error) {
	finally, err := repo.BeginOrContinueTx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	prev, err := repo.FindActByCode(ctx, target.Code())
	if repository.IsNotFound(err) {
		return repo.insertNewAct(repo.tx, ctx, target)
	} else if err != nil {
		return nil, err
	}
	return repo.updateVersions(ctx, prev, target)
}

func (repo *EntActRepository) insertNewAct(tx *ent.Tx, ctx context.Context, target *act.Act) (*act.Act, error) {
	eat, err := entActTypeFromActType(target.Type())
	if err != nil {
		return nil, err
	}
	eact, err := repo.tx.Act.Create().
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
	states := make([]*ent.ActVersionCreate, len(vers))
	for i, v := range vers {
		states[i] = repo.tx.ActVersion.Create().
			SetCode(string(v.Code())).SetVersion(uint(v.Version())).
			SetText(string(v.Text())).SetCreatedAt(time.Time(v.CreatedAt())).
			SetAct(eact).SetActCode(eact.Code)
	}
	evs, err := repo.tx.ActVersion.CreateBulk(states...).Save(ctx)
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
	states := make([]*ent.ActImageCreate, len(imgs))
	for i, img := range imgs {
		states[i] = repo.tx.ActImage.Create().
			SetActVersionCode(ver.Code).SetOrder(uint(i) + 1).SetSrc(img)
	}
	return repo.tx.ActImage.CreateBulk(states...).Save(ctx)
}

func (repo *EntActRepository) updateVersions(ctx context.Context, prev *act.Act, target *act.Act) (*act.Act, error) {
	return nil, nil
}
