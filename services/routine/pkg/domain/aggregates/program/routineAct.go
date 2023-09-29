package program

import "routine/pkg/domain/aggregates/act"

func RoutineActFrom(
	dailyRoutine DailyRoutineCode,
	order RoutineActOrder,
	actVersion act.ActVersionCode,
	stage RoutineActStage,
	repsOrMeters RepsOrMeters,
	ratioOrSecs RatioOrSecs,
) *RoutineAct {
	return &RoutineAct{
		dailyRoutine: dailyRoutine,
		order:        order,
		actVersion:   actVersion,
		stage:        stage,
		repsOrMeters: repsOrMeters,
		ratioOrSecs:  ratioOrSecs,
	}
}

type RoutineActOrder uint
type RepsOrMeters uint
type RatioOrSecs float64

type RoutineAct struct {
	dailyRoutine DailyRoutineCode
	order        RoutineActOrder
	actVersion   act.ActVersionCode
	stage        RoutineActStage
	repsOrMeters RepsOrMeters
	ratioOrSecs  RatioOrSecs
}
