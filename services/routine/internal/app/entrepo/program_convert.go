package entrepo

import (
	"fmt"
	"routine/internal/domain"
	"routine/internal/domain/aggregates/act"
	"routine/internal/domain/aggregates/program"
	"routine/internal/domain/aggregates/user"
	"routine/internal/ent"
	eprogram "routine/internal/ent/program"
	"routine/internal/ent/routineact"
)

func ProgramFromEntProgram(ep *ent.Program) (*program.Program, error) {
	switch ep.ProgramType {
	case eprogram.ProgramTypeWeekly:
		author := user.UserFrom(user.UserId(ep.Author))
		if ep.Edges.ProgramVersions == nil {
			return nil, fmt.Errorf("program versions not found")
		}
		versions, err := VersionsFromEnt(ep.Edges.ProgramVersions)
		if err != nil {
			return nil, err
		}
		p, err := program.WeeklyProgramFrom(
			program.ProgramCode(ep.Code),
			program.ProgramTitle(ep.Title),
			*author,
			domain.CreatedAt(ep.CreatedAt),
			(*program.ProgramVersionCode)(ep.VersionDerivedFrom),
			versions,
		)
		if err != nil {
			return nil, err
		}
		return p, nil
	default:
		return nil, fmt.Errorf("unknown program type: %s", ep.ProgramType)
	}
}

func VersionsFromEnt(epvs []*ent.ProgramVersion) (pvs []*program.ProgramVersion, err error) {
	pvs = make([]*program.ProgramVersion, len(epvs))
	for i, epv := range epvs {
		pv, err := VersionFromEnt(epv)
		if err != nil {
			return nil, err
		}
		pvs[i] = pv
	}
	return pvs, nil
}

func VersionFromEnt(epv *ent.ProgramVersion) (*program.ProgramVersion, error) {
	if epv.Edges.DailyRoutines == nil {
		return nil, fmt.Errorf("daily routines not found")
	}
	if epv.Edges.ProgramImages == nil {
		return nil, fmt.Errorf("program images not found")
	}
	drs, err := DailyRoutinesFromEnt(epv.Edges.DailyRoutines)
	if err != nil {
		return nil, err
	}

	pvimgs, err := ImagesFromProgramImages(epv.Edges.ProgramImages)
	if err != nil {
		return nil, err
	}

	pv, err := program.ProgramVersionFrom(
		program.ProgramVersionCode(epv.Code),
		program.ProgramCode(epv.ProgramCode),
		program.ProgramVersionNumber(epv.Version),
		domain.CreatedAt(epv.CreatedAt),
		pvimgs,
		program.ProgramText(epv.Text),
		drs,
	)
	if err != nil {
		return nil, err
	}
	return pv, nil
}

func ImagesFromProgramImages(epis []*ent.ProgramImage) (program.ProgramImageSrcs, error) {
	imgs := make(program.ProgramImageSrcs, len(epis))
	for i, epi := range epis {
		if epi.Edges.Image == nil {
			return nil, fmt.Errorf("image not found")
		}
		imgs[i] = epi.Edges.Image.Src
	}
	return imgs, nil
}

func DailyRoutinesFromEnt(edrs []*ent.DailyRoutine) ([]*program.DailyRoutine, error) {
	drs := make([]*program.DailyRoutine, len(edrs))
	for i, edr := range edrs {
		dr, err := DailyRoutineFromEnt(edr)
		if err != nil {
			return nil, err
		}
		drs[i] = dr
	}
	return drs, nil
}

func DailyRoutineFromEnt(edr *ent.DailyRoutine) (*program.DailyRoutine, error) {
	if edr.Edges.RoutineActs == nil {
		return nil, fmt.Errorf("routine acts not found")
	}
	ras, err := RoutineActsFromEnt(edr.Code, edr.Edges.RoutineActs)
	if err != nil {
		return nil, err
	}
	dr, err := program.DailyRoutineFrom(
		program.DailyRoutineCode(edr.Code),
		program.ProgramVersionCode(edr.ProgramVersionCode),
		program.DailyRoutineDay(edr.Day),
		ras,
	)
	if err != nil {
		return nil, err
	}
	return dr, nil
}

func RoutineActsFromEnt(drcode string, eras []*ent.RoutineAct) ([]*program.RoutineAct, error) {
	ras := make([]*program.RoutineAct, len(eras))
	for i, era := range eras {
		ra, err := RoutineActFromEnt(drcode, era)
		if err != nil {
			return nil, err
		}
		ras[i] = ra
	}
	return ras, nil
}

func RoutineActFromEnt(drcode string, era *ent.RoutineAct) (*program.RoutineAct, error) {
	stage, err := StageFromEntStage(era.Stage)
	if err != nil {
		return nil, err
	}
	ra := program.RoutineActFrom(
		program.DailyRoutineCode(drcode),
		program.RoutineActOrder(era.Order),
		act.ActVersionCode(era.ActVersion),
		stage,
		program.RepsOrMeters(era.RepsOrMeters),
		program.RatioOrSecs(era.RatioOrSecs),
	)
	if err != nil {
		return nil, err
	}
	return ra, nil
}

func StageFromEntStage(es routineact.Stage) (program.RoutineActStage, error) {
	switch es {
	case program.WARMUP:
		return program.WarmUpStage, nil
	case program.MAIN:
		return program.MainStage, nil
	case program.COOLDOWN:
		return program.CoolDownStage, nil
	default:
		return program.RoutineActStage{}, fmt.Errorf("unknown stage: %s", es)
	}
}

func entProgramTypeFrom(pt string) (eprogram.ProgramType, error) {
	switch pt {
	case program.WEEKLY:
		return eprogram.ProgramTypeWeekly, nil
	default:
		return "", fmt.Errorf("unknown program type: %s", pt)
	}
}
