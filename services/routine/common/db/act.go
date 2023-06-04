package db

import (
	"context"
	"routine/common/dto"
	"routine/ent"
	"routine/ent/act"
)

// QueryActsByName queries acts by name and returns 5 acts skipping given number of acts
func QueryActsByName(dbClient *ent.Client, c context.Context, name string, skip int) (acts []*ent.Act, err error) {
	acts, err = dbClient.Act.Query().
		Where(act.NameContains(name)).
		Offset(skip).
		Limit(5).
		All(c)
	if ent.IsNotFound(err) {
		acts = []*ent.Act{}
	} else if err != nil {
		return nil, err
	}
	return acts, nil
}

// CreateAct creates act and returns created act's ID
func CreateAct(dbClient *ent.Client, c context.Context, newActDto *dto.CreateActDto) (aid uint64, err error) {

	// first query tags and create tags if not exists
	tags, err := QueryAndCreateTags(dbClient, c, newActDto.Tags)
	if err != nil {
		return 0, err
	}

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
		AddTagIDs(tags...).
		Save(c)
	if err != nil {
		return 0, err
	}

	return newAct.ID, nil
}
