package programqry

import (
	"context"
	"routine/internal/app/dto"
	"routine/internal/app/entrepo"
	"routine/internal/domain/aggregates/program"
	"routine/internal/repository"
)

func newProgramQueryService() *programQueryService {
	return &programQueryService{programRepo: repository.NewProgramRepository(entrepo.NewEntProgramRepository())}
}

type programQueryService struct {
	programRepo *repository.ProgramRepository
}

func (ps *programQueryService) findProgramByCode(ctx context.Context, code string) (qpDto *dto.QueryProgramDto, err error) {
	qp, err := ps.programRepo.FindProgramByCode(ctx, program.ProgramCode(code))
	if err != nil {
		return nil, err
	}
	return dto.QueryProgramDtoFrom(qp), nil
}

func (ps *programQueryService) findProgramsByTitle(ctx context.Context, title string) ([]*dto.QueryProgramDto, error) {
	qps, err := ps.programRepo.FindProgramsByTitle(ctx, program.ProgramTitle(title))
	if err != nil {
		return nil, err
	}
	qpdtos := make([]*dto.QueryProgramDto, len(qps))
	for i, qp := range qps {
		qpdtos[i] = dto.QueryProgramDtoFrom(qp)
	}
	return qpdtos, nil
}
