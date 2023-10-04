package entrepo

import (
	"context"
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

func (repo *EntActRepository) FindByCode(ctx context.Context, code act.ActCode) (fAct *act.Act, err error) {
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
	prev, err := repo.FindByCode(ctx, target.Code())
	if ent.IsNotFound(err) {
		return repo.insertNewAct(repo.tx, ctx, target)
	} else if err != nil {
		return nil, err
	}
	return repo.updateAct(ctx, prev, target)
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
	vs, err := repo.tx.ActVersion.CreateBulk(
		repo.tx.ActVersion.Create().
			SetCode(string(target.LatestVersion().Code())).SetVersion(1).
			SetText(string(target.LatestVersion().Text())).SetCreatedAt(time.Time(target.LatestVersion().CreatedAt())).
			SetAct(eact).SetActCode(eact.Code),
	).Save(ctx)
	if err != nil {
		return nil, err
	}
	imgs, err := repo.tx.ActImage.CreateBulk(
		repo.imgsToEntCreateStates(vs[0].Code, target.LatestVersion().ImageSrcs())...,
	).Save(ctx)
	if err != nil {
		return nil, err
	}
	vs[0].Edges.ActImages = imgs
	eact.Edges.ActVersions = vs
	return repo.actFromEntAct(ctx, eact)
}

func (repo *EntActRepository) imgsToEntCreateStates(verCode string, imgs act.ActImageSrcs) []*ent.ActImageCreate {
	states := make([]*ent.ActImageCreate, len(imgs))
	for i, img := range imgs {
		states[i] = repo.tx.ActImage.Create().
			SetActVersionCode(verCode).SetOrder(uint(i) + 1).SetSrc(img)
	}
	return states
}

func (repo *EntActRepository) updateAct(ctx context.Context, prev *act.Act, target *act.Act) (*act.Act, error) {
	return nil, nil
}
