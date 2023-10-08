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
	return ProgramFromEntProgram(ep)
}

func (repo *EntProgramRepository) Save(ctx context.Context, target *program.Program) (saved *program.Program, err error) {
	return nil, nil
}
