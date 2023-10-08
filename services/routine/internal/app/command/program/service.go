package programcmd

import (
	"routine/internal/app/dto"
	"routine/internal/app/entrepo"
	"routine/internal/repository"
)

func newProgramCommandService() *programCommandService {
	return &programCommandService{programRepo: repository.NewProgramRepository(entrepo.NewEntProgramRepository())}
}

type programCommandService struct {
	programRepo *repository.ProgramRepository
}

func (ps *programCommandService) createProgram(cpDto dto.CreateProgramServiceDto) (qpDto *dto.QueryProgramDto, err error) {
	return nil, nil
}
