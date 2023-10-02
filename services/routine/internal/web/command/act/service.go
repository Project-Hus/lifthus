package act

import (
	"fmt"
	"routine/internal/domain/aggregates/act"
	"routine/internal/domain/aggregates/user"
	"routine/internal/web/dto"
)

type actService struct {
}

func (as *actService) createAct(dto dto.CreateActDto) (*dto.QueryActDto, error) {
	actType, err := act.MapActType(dto.ActType)
	if err != nil {
		return nil, err
	}
	author := user.UserFrom(user.UserId(dto.Author))
	newAct, err := act.CreateAct(*actType, act.ActName(dto.Name), *author, dto.ImageSrcs, act.ActText(dto.Text))
	if err != nil {
		return nil, err
	}
	fmt.Println(newAct)
	return nil, nil
}

func (as *actService) upgradeAct(dto dto.UpgradeActDto) (*dto.QueryActDto, error) {
	return nil, nil
}
