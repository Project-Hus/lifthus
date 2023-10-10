package programcmd

import (
	"context"
	"routine/internal/app/dto"
	"routine/internal/app/entrepo"
	"routine/internal/domain/aggregates/program"
	"routine/internal/domain/aggregates/user"
	"routine/internal/repository"
)

func newProgramCommandService() *programCommandService {
	return &programCommandService{programRepo: repository.NewProgramRepository(entrepo.NewEntProgramRepository())}
}

type programCommandService struct {
	programRepo *repository.ProgramRepository
}

func (ps *programCommandService) createWeeklyProgram(ctx context.Context, cpDto dto.CreateProgramServiceDto) (qpDto *dto.QueryProgramDto, err error) {
	author := user.UserFrom(user.UserId(cpDto.Author))
	drs, err := dailyRoutinesFrom(cpDto.DailyRoutines)
	if err != nil {
		return nil, err
	}
	program, err := program.CreateWeeklyProgram(
		program.ProgramTitle(cpDto.Title),
		*author,
		(*program.ProgramVersionCode)(cpDto.DerivedFrom),
		cpDto.ImageSrcs,
		program.ProgramText(cpDto.Text),
		drs,
	)
	if err != nil {
		return nil, err
	}
	program, err = ps.programRepo.Save(ctx, program)
	if err != nil {
		return nil, err
	}
	return dto.QueryProgramDtoFrom(program), nil
}
