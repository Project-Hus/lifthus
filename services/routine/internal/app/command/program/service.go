package programcmd

import (
	"routine/internal/app/entrepo"
	"routine/internal/repository"
)

func newProgramCommandService() *programCommandService {
	return &programCommandService{programRepo: repository.NewProgramRepository(entrepo.NewEntProgramRepository())}
}

type programCommandService struct {
	programRepo *repository.ProgramRepository
}
