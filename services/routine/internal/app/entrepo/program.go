package entrepo

import (
	"context"
	"routine/internal/domain/aggregates/program"
	"routine/internal/ent"
	eprogram "routine/internal/ent/program"
	"routine/internal/ent/programrelease"
	"routine/internal/ent/routine"
	"routine/internal/ent/routineact"
	"routine/internal/ent/s3programimage"
	"routine/internal/repository"
)

func NewEntProgramRepository() *EntProgramRepository {
	return &EntProgramRepository{EntRepository: NewEntRepository()}
}

type EntProgramRepository struct {
	*EntRepository
}

func (repo *EntProgramRepository) FindProgramByCode(ctx context.Context, code program.ProgramCode) (fProgram *program.Program, err error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	ep, err := tx.Program.Query().Where(eprogram.Code(string(code))).
		WithProgramReleases(
			func(q *ent.ProgramReleaseQuery) {
				q.Order(ent.Asc(programrelease.FieldVersion))
				q.WithS3ProgramImages(
					func(q *ent.S3ProgramImageQuery) {
						q.Order(ent.Asc(s3programimage.FieldOrder))
						q.WithS3Image()
					},
				)
				q.WithRoutines(
					func(q *ent.RoutineQuery) {
						q.Order(ent.Asc(routine.FieldDay))
						q.WithRoutineActs(
							func(q *ent.RoutineActQuery) {
								q.Order(ent.Asc(routineact.FieldOrder))
							},
						)
					},
				)
			},
		).
		First(ctx)
	if ent.IsNotFound(err) {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return ProgramFromEntProgram(ep)
}

func (repo *EntProgramRepository) FindProgramsByTitle(ctx context.Context, title program.ProgramTitle) ([]*program.Program, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	eps, err := tx.Program.Query().Where(eprogram.TitleContains(string(title))).
		WithProgramReleases(
			func(q *ent.ProgramReleaseQuery) {
				q.Order(ent.Asc(programrelease.FieldVersion))
				q.WithS3ProgramImages(
					func(q *ent.S3ProgramImageQuery) {
						q.Order(ent.Asc(s3programimage.FieldOrder))
						q.WithS3Image()
					},
				)
				q.WithRoutines(
					func(q *ent.RoutineQuery) {
						q.Order(ent.Asc(routine.FieldDay))
						q.WithRoutineActs(
							func(q *ent.RoutineActQuery) {
								q.Order(ent.Asc(routineact.FieldOrder))
							},
						)
					},
				)
			},
		).All(ctx)
	if err != nil {
		return nil, err
	}
	dps := make([]*program.Program, len(eps))
	for i, ep := range eps {
		dps[i], err = ProgramFromEntProgram(ep)
		if err != nil {
			return nil, err
		}
	}
	return dps, nil
}

func (repo *EntProgramRepository) Save(ctx context.Context, target *program.Program) (saved *program.Program, err error) {
	finally, err := repo.BeginOrContinueTx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	existing, err := repo.FindProgramByCode(ctx, target.Code())
	if repository.IsNotFound(err) {
		return repo.insertNewProgram(ctx, target)
	} else if err != nil {
		return nil, err
	}
	return repo.updateProgram(ctx, existing, target)
}

func (repo *EntProgramRepository) updateProgram(ctx context.Context, existing *program.Program, target *program.Program) (*program.Program, error) {
	return nil, nil
}
