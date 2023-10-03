package repository

import (
	"context"
	"routine/internal/domain/aggregates/act"
)

func NewActRepository(actRepo actRepository) *ActRepository {
	return &ActRepository{repo: actRepo}
}

type actRepository interface {
	Save(ctx context.Context, act *act.Act) (*act.Act, error)
	FindByCode(ctx context.Context, code act.ActCode) (*act.Act, error)

	Commit() error
	Rollback(err error) error
}

type ActRepository struct {
	repo actRepository
}

func (ar *ActRepository) Save(ctx context.Context, act *act.Act) (*act.Act, error) {
	return ar.repo.Save(ctx, act)
}

func (ar *ActRepository) FindByCode(ctx context.Context, code act.ActCode) (*act.Act, error) {
	return ar.repo.FindByCode(ctx, code)
}
