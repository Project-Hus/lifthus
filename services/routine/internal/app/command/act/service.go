package actcmd

import (
	"context"
	"routine/internal/app/dto"
	"routine/internal/app/entrepo"
	"routine/internal/domain/aggregates/act"
	"routine/internal/domain/aggregates/user"
	"routine/internal/repository"
)

func newActCommandService() *actCommandService {
	return &actCommandService{actRepo: repository.NewActRepository(entrepo.NewEntActRepository())}
}

type actCommandService struct {
	actRepo *repository.ActRepository
}

func (as *actCommandService) createAct(ctx context.Context, caDto dto.CreateActServiceDto) (qaDto *dto.QueryActDto, err error) {
	finally, err := as.actRepo.BeginOrContinueTx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	actType, err := act.MapActType(caDto.ActType)
	if err != nil {
		return nil, err
	}
	author := user.UserFrom(user.UserId(caDto.Author))
	newAct, err := act.CreateAct(*actType, act.ActName(caDto.Name), *author, caDto.ImageSrcs, act.ActText(caDto.Text))
	if err != nil {
		return nil, err
	}
	act, err := as.actRepo.Save(ctx, newAct)
	if err != nil {
		return nil, err
	}
	return dto.QueryActDtoFrom(act), nil
}
