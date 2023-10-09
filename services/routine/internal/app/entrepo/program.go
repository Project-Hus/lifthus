package entrepo

import (
	"context"
	"routine/internal/domain/aggregates/program"
	"routine/internal/ent"
	"routine/internal/ent/dailyroutine"
	eprogram "routine/internal/ent/program"
	"routine/internal/ent/programimage"
	eprogramversion "routine/internal/ent/programversion"
	"routine/internal/ent/routineact"
	"routine/internal/repository"
	"time"
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
		WithProgramVersions(
			func(q *ent.ProgramVersionQuery) {
				q.Order(ent.Asc(eprogramversion.FieldVersion))
				q.WithProgramImages(
					func(q *ent.ProgramImageQuery) {
						q.Order(ent.Asc(programimage.FieldOrder))
						q.WithImage()
					},
				)
				q.WithDailyRoutines(
					func(q *ent.DailyRoutineQuery) {
						q.Order(ent.Asc(dailyroutine.FieldDay))
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

func (repo *EntProgramRepository) insertNewProgram(ctx context.Context, np *program.Program) (*program.Program, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	ptype, err := entProgramTypeFrom(np.ProgramType().Type())
	if err != nil {
		return nil, err
	}
	_, err = tx.Program.Create().
		SetCode(string(np.Code())).
		SetProgramType(ptype).
		SetTitle(string(np.Title())).
		SetAuthor(uint64(np.Author())).
		SetCreatedAt(time.Time(np.CreatedAt())).
		SetNillableVersionDerivedFrom((*string)(np.DerivedFrom())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *EntProgramRepository) updateProgram(ctx context.Context, existing *program.Program, target *program.Program) (*program.Program, error) {
	return nil, nil
}
