package programcmd

import (
	"routine/internal/app/dto"
	"routine/internal/domain/aggregates/act"
	"routine/internal/domain/aggregates/program"
)

func dailyRoutinesFrom(drdtos []dto.CreateProgramServiceDailyRoutineDto) ([]*program.DailyRoutine, error) {
	drs := make([]*program.DailyRoutine, len(drdtos))
	for i, drdto := range drdtos {
		dr, err := dailyRoutineFrom(drdto)
		if err != nil {
			return nil, err
		}
		drs[i] = dr
	}
	return drs, nil
}

func dailyRoutineFrom(drdto dto.CreateProgramServiceDailyRoutineDto) (*program.DailyRoutine, error) {
	ras, err := routineActsFrom(drdto.RoutineActs)
	if err != nil {
		return nil, err
	}
	return program.CreateDailyRoutineWithoutProgramVersion(
		program.DailyRoutineDay(drdto.Day),
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
	return program.CreateRoutineActWithoutDailyRoutine(
		program.RoutineActOrder(radto.Order),
		act.ActVersionCode(radto.ActVersion),
		stage,
		program.RepsOrMeters(radto.RepsOrMeters),
		program.RatioOrSecs(radto.RatioOrSecs),
	), nil
}
