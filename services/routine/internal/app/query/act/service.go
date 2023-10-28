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
	dact, err := as.actRepo.FindActByCode(ctx, act.ActCode(code))
	if err != nil {
		return nil, err
	}
	return dto.QueryActDtoFrom(dact), nil
}

func (as *actQueryService) queryActsByName(ctx context.Context, actName string) (qaDtos []*dto.QueryActDto, err error) {
	dacts, err := as.actRepo.FindActsByName(ctx, actName)
	if err != nil {
		return nil, err
	}
	qaDtos = make([]*dto.QueryActDto, len(dacts))
	for i, dact := range dacts {
		qaDtos[i] = dto.QueryActDtoFrom(dact)
	}
	return qaDtos, nil
}
