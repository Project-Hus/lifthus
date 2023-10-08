package repository

import (
	"context"
	"routine/internal/domain/aggregates/program"
)

func NewProgramRepository(programRepo programRepository) *ProgramRepository {
	return &ProgramRepository{repo: programRepo}
}

type programRepository interface {
	Save(ctx context.Context, p *program.Program) (*program.Program, error)
	FindProgramByCode(ctx context.Context, code program.ProgramCode) (*program.Program, error)

	BeginOrContinueTx(ctx context.Context) (func(*error), error)
	Commit() error
	Rollback(err error) error
}

type ProgramRepository struct {
	repo programRepository
}

func (pr *ProgramRepository) BeginOrContinueTx(ctx context.Context) (func(*error), error) {
	return pr.repo.BeginOrContinueTx(ctx)
}

func (pr *ProgramRepository) Save(ctx context.Context, p *program.Program) (*program.Program, error) {
	return pr.repo.Save(ctx, p)
}

func (pr *ProgramRepository) FindProgramByCode(ctx context.Context, code program.ProgramCode) (*program.Program, error) {
	return pr.repo.FindProgramByCode(ctx, code)
}
