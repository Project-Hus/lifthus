package db

import (
	"context"
	"routine/common/dto"
	"routine/ent"
	"routine/ent/act"
)

func CreateAct(dbClient *ent.Client, c context.Context, newActDto *dto.CreateActDto) (aid uint, err error) {
	newAct, err := dbClient.Act.Create().
		SetName(newActDto.Name).
		SetType(act.Type(newActDto.Type)).
		SetAuthor(newActDto.Author).
		SetNillableImage(newActDto.Image).
		SetNillableDescription(newActDto.Description).
		SetWeight(newActDto.Weight).
		SetBodyweight(newActDto.Bodyweight).
		SetCardio(newActDto.Cardio).
		SetUpper(newActDto.Upper).
		SetLower(newActDto.Lower).
		SetArms(newActDto.Arms).
		SetShoulders(newActDto.Shoulders).
		SetChest(newActDto.Chest).
		SetCore(newActDto.Core).
		SetUpperBack(newActDto.UpperBack).
		SetLowerBack(newActDto.LowerBack).
		SetGlute(newActDto.Glute).
		SetLegsFront(newActDto.LegsFront).
		SetLegsBack(newActDto.LegsBack).
		SetEtc(newActDto.Etc).
		Save(c)
	if err != nil {
		return 0, err
	}

	return uint(newAct.ID), nil
}
