package entrepo

import (
	"context"
	"routine/internal/domain"
	"routine/internal/domain/aggregates/act"
	"routine/internal/domain/aggregates/user"
	"routine/internal/ent"
	eact "routine/internal/ent/act"
	eacti "routine/internal/ent/actimage"
	eactv "routine/internal/ent/actversion"
	"routine/internal/repository"
	"routine/pkg/db"
)

func NewEntActRepository() *EntActRepository {
	return &EntActRepository{c: db.EntClient()}
}

type EntActRepository struct {
	c  *ent.Client
	tx *ent.Tx
}

func (repo *EntActRepository) getOrGenTx(ctx context.Context) (*ent.Tx, error) {
	if repo.tx != nil {
		return repo.tx, nil
	}
	tx, err := repo.c.Tx(ctx)
	if err != nil {
		return nil, err
	}
	repo.tx = tx
	return repo.tx, nil
}

func (repo *EntActRepository) Save(ctx context.Context, target *act.Act) (*act.Act, error) {
	tx, err := repo.getOrGenTx(ctx)
	if err != nil {
		return nil, err
	}
	exsisting, err := tx.Act.Query().
		Where(eact.Code(string(target.Code()))).
		WithActVersions(
			func(q *ent.ActVersionQuery) {

			},
		).
		First(ctx)
	if ent.IsNotFound(err) {
		return repo.insertNewAct(ctx, target)
	} else if err != nil {
		return nil, err
	}
	exsisting.Update()
	return repo.upgradeAct(ctx, target)
}

func (repo *EntActRepository) insertNewAct(ctx context.Context, target *act.Act) (*act.Act, error) {
	return nil, nil
}

func (repo *EntActRepository) upgradeAct(ctx context.Context, target *act.Act) (*act.Act, error) {
	return nil, nil
}

func (repo *EntActRepository) FindByCode(ctx context.Context, code act.ActCode) (*act.Act, error) {
	tx, err := repo.getOrGenTx(ctx)
	if err != nil {
		return nil, err
	}
	a, err := tx.Act.Query().
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

func (repo *EntActRepository) actFromEntAct(ctx context.Context, ea *ent.Act) (*act.Act, error) {
	actType, err := act.MapActType(string(ea.ActType))
	if err != nil {
		return nil, err
	}
	versions := ea.Edges.ActVersions
	vs := repo.actVersionsFromEntVersions(ctx, versions)
	return act.ActFrom(act.ActCode(ea.Code), *actType, act.ActName(ea.Name), user.UserId(ea.Author), domain.CreatedAt(ea.CreatedAt), vs)
}

func (repo *EntActRepository) actVersionsFromEntVersions(ctx context.Context, evs []*ent.ActVersion) []*act.ActVersion {
	versions := make([]*act.ActVersion, len(evs))
	for i, ev := range evs {
		imgs := ev.Edges.ActImages
		imgSrcs := imgSrcsFromEntImgs(imgs)
		v := act.ActVersionFrom(
			act.ActVersionCode(ev.ActCode),
			act.ActVersionNumber(ev.Version),
			imgSrcs,
			act.ActText(ev.Text),
			domain.CreatedAt(ev.CreatedAt),
		)
		versions[i] = v
	}
	return versions
}

func imgSrcsFromEntImgs(eis []*ent.ActImage) act.ActImageSrcs {
	imgSrcs := make(act.ActImageSrcs, len(eis))
	for i, ei := range eis {
		imgSrcs[i] = ei.Src
	}
	return imgSrcs
}
