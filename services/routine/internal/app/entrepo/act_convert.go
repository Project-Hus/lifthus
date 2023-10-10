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
	versions := ea.Edges.ActVersions
	vs := repo.actVersionsFromEntVersions(ctx, versions)
	return act.ActFrom(act.ActCode(ea.Code), *actType, act.ActName(ea.Name), user.UserId(ea.Author), domain.CreatedAt(ea.CreatedAt), vs)
}

func (repo *EntActRepository) actVersionsFromEntVersions(ctx context.Context, evs []*ent.ActVersion) []*act.ActVersion {
	versions := make([]*act.ActVersion, len(evs))
	for i, ev := range evs {
		imgs := ev.Edges.ActImages
		imgSrcs := imgSrcsFromEntImgs(imgs)
		v := act.ActVersionFrom(
			act.ActVersionCode(ev.Code),
			act.ActVersionNumber(ev.Version),
			imgSrcs,
			act.ActText(ev.Text),
			domain.CreatedAt(ev.CreatedAt),
		)
		versions[i] = v
	}
	return versions
}

func imgSrcsFromEntImgs(eis []*ent.ActImage) act.ActImageSrcs {
	imgSrcs := make(act.ActImageSrcs, len(eis))
	for i, ei := range eis {
		imgSrcs[i] = ei.Edges.Image.Src
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
