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

func (ps *programCommandService) createProgram(ctx context.Context, cpDto dto.CreateProgramServiceDto) (qpDto *dto.QueryProgramDto, err error) {
	ptype, err := program.MapProgramType(cpDto.ProgramType)
	if err != nil {
		return nil, err
	}
	author := user.UserFrom(user.UserId(cpDto.Author))
	rs, err := routinesFrom(cpDto.Routines)
	if err != nil {
		return nil, err
	}
	v1, err := program.CreateProgramRelease(
		1,
		cpDto.ImageSrcs,
		program.ProgramText(cpDto.Text),
		rs,
	)
	if err != nil {
		return nil, err
	}
	var pv *program.ParentProgramVersion
	if cpDto.ParentProgram != nil {
		pv = &program.ParentProgramVersion{
			ProgramCode:          program.ProgramCode(*cpDto.ParentProgram),
			ProgramVersionNumber: program.ProgramVersionNumber(*cpDto.ParentVersion),
		}
	}
	program, err := program.CreateProgram(
		*ptype,
		program.ProgramTitle(cpDto.Title),
		author.Id(),
		pv,
		*v1,
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
