package program

import "routine/pkg/domain/aggregates/act"

type RoutineActOrder uint
type RepsOrMeters uint
type RatioOrSecs float64

type RoutineAct struct {
	dailyRoutine DailyRoutineCode
	order        RoutineActOrder
	act          act.ActCode
	actVersion   act.ActVersionNumber
	stage        RoutineActStage
	repsOrMeters RepsOrMeters
	ratioOrSecs  RatioOrSecs
}
