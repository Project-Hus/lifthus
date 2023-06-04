package db

import (
	"context"
	"routine/common/dto"
	"routine/ent"
)

func CreateWeeklyProgramRec(
	dbClient *ent.Client,
	c context.Context,
	newPRec dto.CreateWeeklyProgramRecDto,
) (rid uint64, err error) {
	return 0, nil
}
