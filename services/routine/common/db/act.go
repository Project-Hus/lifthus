package db

import (
	"context"
	"routine/common/dto"
	"routine/ent"
)

func CreateAct(dbClient *ent.Client, c context.Context, newAct *dto.CreateActDto) (aid uint, err error) {
	return 0, nil
}
