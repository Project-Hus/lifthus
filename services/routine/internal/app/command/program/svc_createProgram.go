package programcmd

import (
	"routine/internal/app/dto"
	"routine/internal/domain/aggregates/act"
	"routine/internal/domain/aggregates/program"
)

func routinesFrom(drdtos []dto.CreateProgramServiceRoutineDto) ([]*program.Routine, error) {
	drs := make([]*program.Routine, len(drdtos))
	for i, drdto := range drdtos {
		dr, err := routineFrom(drdto)
		if err != nil {
			return nil, err
		}
		drs[i] = dr
	}
	return drs, nil
}

func routineFrom(drdto dto.CreateProgramServiceRoutineDto) (*program.Routine, error) {
	ras, err := routineActsFrom(drdto.RoutineActs)
	if err != nil {
		return nil, err
	}
	return program.CreateRoutine(
		program.RoutineDay(drdto.Day),
		ras,
	)
}

func routineActsFrom(radtos []dto.CreateProgramServiceRoutineActDto) ([]*program.RoutineAct, error) {
	ras := make([]*program.RoutineAct, len(radtos))
	for i, radto := range radtos {
		ra, err := routineActFrom(radto)
		if err != nil {
			return nil, err
		}
		ras[i] = ra
	}
	return ras, nil
}

func routineActFrom(radto dto.CreateProgramServiceRoutineActDto) (*program.RoutineAct, error) {
	stage, err := program.MapRoutineActStage(radto.Stage)
	if err != nil {
		return nil, err
	}
	return program.CreateRoutineAct(
		program.RoutineActOrder(radto.Order),
		act.ActCode(radto.ActCode),
		stage,
		program.RepsOrMeters(radto.RepsOrMeters),
		program.RatioOrSecs(radto.RatioOrSecs),
	), nil
}
