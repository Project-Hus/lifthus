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
	FindActByCode(ctx context.Context, code act.ActCode) (*act.Act, error)

	BeginOrContinueTx(ctx context.Context) (func(*error), error)
	Commit() error
	Rollback(err error) error
}

type ActRepository struct {
	repo actRepository
}

func (ar *ActRepository) BeginOrContinueTx(ctx context.Context) (func(*error), error) {
	return ar.repo.BeginOrContinueTx(ctx)
}

func (ar *ActRepository) Save(ctx context.Context, act *act.Act) (*act.Act, error) {
	return ar.repo.Save(ctx, act)
}

func (ar *ActRepository) FindActByCode(ctx context.Context, code act.ActCode) (*act.Act, error) {
	return ar.repo.FindActByCode(ctx, code)
}
