package entrepo

import (
	"context"
	"routine/internal/domain/aggregates/act"
	"routine/internal/ent"
	eact "routine/internal/ent/act"
	"routine/internal/ent/s3actimage"
	"routine/internal/ent/s3image"
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
	ea, err := tx.Act.Query().
		Where(eact.Code(string(code))).
		WithS3ActImages(
			func(q *ent.S3ActImageQuery) {
				q.Order(ent.Asc(s3actimage.FieldOrder))
				q.WithS3Image()
			},
		).
		First(ctx)
	if ent.IsNotFound(err) {
		return nil, repository.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return repo.actFromEntAct(ctx, ea)
}

func (repo *EntActRepository) FindActsByName(ctx context.Context, actName string) (fActs []*act.Act, err error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	eacts, err := tx.Act.Query().Where(eact.NameContains(actName)).
		WithS3ActImages(
			func(q *ent.S3ActImageQuery) {
				q.Order(ent.Asc(s3actimage.FieldOrder))
				q.WithS3Image()
			},
		).
		All(ctx)
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
	return repo.updateAct(ctx, prev, target)
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
		SetAuthor(int64(target.Author())).
		SetActType(eat).
		SetName(string(target.Name())).
		SetText(string(target.Text())).
		SetCreatedAt(time.Time(target.CreatedAt())).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	imgs, err := repo.createImgBulk(ctx, eact, target.ImageSrcs())

	eact.Edges.S3ActImages = imgs
	return repo.actFromEntAct(ctx, eact)
}

func (repo *EntActRepository) createImgBulk(ctx context.Context, eact *ent.Act, imgSrcs act.ActImageSrcs) ([]*ent.S3ActImage, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	states := make([]*ent.S3ActImageCreate, len(imgSrcs))
	for i, src := range imgSrcs {
		eimg, err := tx.S3Image.Query().Where(s3image.Src(src)).First(ctx)
		if err != nil {
			return nil, err
		}
		states[i] = tx.S3ActImage.Create().SetOrder(int(i) + 1).SetS3Image(eimg).SetAct(eact)
	}
	return tx.S3ActImage.CreateBulk(states...).Save(ctx)
}

func (repo *EntActRepository) updateAct(ctx context.Context, prev *act.Act, cur *act.Act) (*act.Act, error) {
	return nil, nil
}
