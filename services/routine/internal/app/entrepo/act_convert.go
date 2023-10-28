package entrepo

import (
	"context"
	"routine/internal/domain"
	"routine/internal/domain/aggregates/act"
	"routine/internal/domain/aggregates/user"
	"routine/internal/ent"
	eact "routine/internal/ent/act"
	"routine/internal/repository"
)

func (repo *EntActRepository) actFromEntAct(ctx context.Context, ea *ent.Act) (*act.Act, error) {
	actType, err := act.MapActType(string(ea.ActType))
	if err != nil {
		return nil, err
	}
	return act.ActFrom(
		act.ActCode(ea.Code),
		user.UserId(ea.Author),
		*actType,
		act.ActName(ea.Name),
		act.ActText(ea.Text),
		imgSrcsFromEntImgs(ea.Edges.S3ActImages),
		domain.CreatedAt(ea.CreatedAt),
		ea.Standard,
	), nil
}

func imgSrcsFromEntImgs(eis []*ent.S3ActImage) act.ActImageSrcs {
	imgSrcs := make(act.ActImageSrcs, len(eis))
	for i, ei := range eis {
		imgSrcs[i] = ei.Edges.S3Image.Src
	}
	return imgSrcs
}

func entActTypeFromActType(at act.ActType) (eact.ActType, error) {
	switch at {
	case act.WeightType:
		return eact.ActTypeWeight, nil
	case act.TimeType:
		return eact.ActTypeTime, nil
	case act.SimpleType:
		return eact.ActTypeSimple, nil
	default:
		return "", repository.ErrInvalidActType
	}
}
