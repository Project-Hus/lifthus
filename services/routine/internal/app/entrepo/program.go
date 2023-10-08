package entrepo

import (
	"context"
	"routine/internal/domain/aggregates/program"
)

func NewEntProgramRepository() *EntProgramRepository {
	return &EntProgramRepository{EntRepository: NewEntRepository()}
}

type EntProgramRepository struct {
	*EntRepository
}

func (repo *EntProgramRepository) FindProgramByCode(ctx context.Context, code program.ProgramCode) (fProgram *program.Program, err error) {
	_, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	//ep, err := tx..Query().Where(eprogram.Code(code)).First(ctx)
	return nil, nil
}

func (repo *EntProgramRepository) Save(ctx context.Context, target *program.Program) (saved *program.Program, err error) {
	return nil, nil
}
