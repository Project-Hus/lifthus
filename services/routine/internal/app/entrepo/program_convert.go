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
	programType, err := program.MapProgramType(string(ep.ProgramType))
	if err != nil {
		return nil, err
	}
	if ep.Edges.ProgramReleases == nil {
		return nil, fmt.Errorf("program releases not found")
	}
	rels, err := ReleasesFromEnt(ep.Edges.ProgramReleases)
	if err != nil {
		return nil, err
	}
	var parent *program.ParentProgramVersion
	if ep.ParentProgram != nil {
		parent = &program.ParentProgramVersion{
			ProgramCode:          program.ProgramCode(*ep.ParentProgram),
			ProgramVersionNumber: program.ProgramVersionNumber(*ep.ParentVersion),
		}
	}
	p, err := program.ProgramFrom(
		program.ProgramCode(ep.Code),
		*programType,
		program.ProgramTitle(ep.Title),
		user.UserId(ep.Author),
		domain.CreatedAt(ep.CreatedAt),
		parent,
		rels,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func ReleasesFromEnt(epvs []*ent.ProgramRelease) (pvs []*program.ProgramRelease, err error) {
	pvs = make([]*program.ProgramRelease, len(epvs))
	for i, epv := range epvs {
		pv, err := ReleaseFromEnt(epv)
		if err != nil {
			return nil, err
		}
		pvs[i] = pv
	}
	return pvs, nil
}

func ReleaseFromEnt(epr *ent.ProgramRelease) (*program.ProgramRelease, error) {
	if epr.Edges.Routines == nil {
		return nil, fmt.Errorf("daily routines not found")
	}
	if epr.Edges.S3ProgramImages == nil {
		return nil, fmt.Errorf("program images not found")
	}
	rs, err := RoutinesFromEnt(epr.Edges.Routines)
	if err != nil {
		return nil, err
	}

	rimgs, err := ImageSrcsFromEnt(epr.Edges.S3ProgramImages)
	if err != nil {
		return nil, err
	}

	pv, err := program.ProgramReleaseFrom(
		program.ProgramVersionNumber(epr.Version),
		domain.CreatedAt(epr.CreatedAt),
		rimgs,
		program.ProgramText(epr.Text),
		rs,
	)
	if err != nil {
		return nil, err
	}
	return pv, nil
}

func ImageSrcsFromEnt(epis []*ent.S3ProgramImage) (program.ProgramImageSrcs, error) {
	imgs := make(program.ProgramImageSrcs, len(epis))
	for i, epi := range epis {
		if epi.Edges.S3Image == nil {
			return nil, fmt.Errorf("image not found")
		}
		imgs[i] = epi.Edges.S3Image.Src
	}
	return imgs, nil
}

func RoutinesFromEnt(edrs []*ent.Routine) ([]*program.Routine, error) {
	drs := make([]*program.Routine, len(edrs))
	for i, edr := range edrs {
		dr, err := RoutineFromEnt(edr)
		if err != nil {
			return nil, err
		}
		drs[i] = dr
	}
	return drs, nil
}

func RoutineFromEnt(edr *ent.Routine) (*program.Routine, error) {
	if edr.Edges.RoutineActs == nil {
		return nil, fmt.Errorf("routine acts not found")
	}
	ras, err := RoutineActsFromEnt(edr.Edges.RoutineActs)
	if err != nil {
		return nil, err
	}
	dr, err := program.RoutineFrom(
		program.RoutineDay(edr.Day),
		ras,
	)
	if err != nil {
		return nil, err
	}
	return dr, nil
}

func RoutineActsFromEnt(eras []*ent.RoutineAct) ([]*program.RoutineAct, error) {
	ras := make([]*program.RoutineAct, len(eras))
	for i, era := range eras {
		ra, err := RoutineActFromEnt(era)
		if err != nil {
			return nil, err
		}
		ras[i] = ra
	}
	return ras, nil
}

func RoutineActFromEnt(era *ent.RoutineAct) (*program.RoutineAct, error) {
	stage, err := StageFromEntStage(era.Stage)
	if err != nil {
		return nil, err
	}
	ra := program.RoutineActFrom(
		program.RoutineActOrder(era.Order),
		act.ActCode(era.ActCode),
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
	case program.DAILY:
		return eprogram.ProgramTypeDaily, nil
	default:
		return "", fmt.Errorf("unknown program type: %s", pt)
	}
}
