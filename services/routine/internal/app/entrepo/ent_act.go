package entrepo

import (
	"context"
	"fmt"
	"routine/internal/domain/aggregates/act"
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

func (repo *EntActRepository) FindByCode(ctx context.Context, code act.ActCode) (*act.Act, error) {
	tx, err := repo.getOrBeginTx(ctx)
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

func (repo *EntActRepository) Save(ctx context.Context, target *act.Act) (*act.Act, error) {
	prev, err := repo.FindByCode(ctx, target.Code())
	if ent.IsNotFound(err) {
		return repo.insertNewAct(ctx, target)
	} else if err != nil {
		return nil, err
	}
	return repo.upgradeAct(ctx, prev, target)
}

func (repo *EntActRepository) insertNewAct(ctx context.Context, target *act.Act) (*act.Act, error) {
	return nil, nil
}

func (repo *EntActRepository) upgradeAct(ctx context.Context, prev *act.Act, target *act.Act) (*act.Act, error) {
	return nil, nil
}

func (repo *EntActRepository) Commit() error {
	if repo.tx == nil {
		return repository.ErrNoTransaction
	}
	return repo.tx.Commit()
}

func (repo *EntActRepository) Rollback(err error) error {
	if repo.tx == nil {
		return repository.ErrNoTransaction
	}
	if rerr := repo.tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
