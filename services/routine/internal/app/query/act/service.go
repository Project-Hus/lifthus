package actQuery

import (
	"context"
	"routine/internal/app/dto"
	"routine/internal/app/entrepo"
	"routine/internal/domain/aggregates/act"
	"routine/internal/repository"
)

func newActQueryService() *actQueryService {
	return &actQueryService{actRepo: repository.NewActRepository(entrepo.NewEntActRepository())}
}

type actQueryService struct {
	actRepo *repository.ActRepository
}

func (as *actQueryService) queryActByCode(ctx context.Context, code string) (qaDto *dto.QueryActDto, err error) {
	eact, err := as.actRepo.FindActByCode(ctx, act.ActCode(code))
	if err != nil {
		return nil, err
	}
	return dto.QueryActDtoFrom(eact), nil
}
