package program

import "routine/pkg/domain/aggregates/act"

func CreateRoutineActWithoutDailyRoutine(
	order RoutineActOrder,
	actVersion act.ActVersionCode,
	stage RoutineActStage,
	repsOrMeters RepsOrMeters,
	ratioOrSecs RatioOrSecs,
) *RoutineAct {
	return &RoutineAct{
		order:        order,
		actVersion:   actVersion,
		stage:        stage,
		repsOrMeters: repsOrMeters,
		ratioOrSecs:  ratioOrSecs,
	}
}

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

type RoutineActs []*RoutineAct

func (ras RoutineActs) IsValid() bool {
	if len(ras) == 0 {
		return false
	}
	for i, ra := range ras {
		if i != int(ra.order)-1 {
			return false
		}
	}
	return true
}

type RoutineAct struct {
	dailyRoutine DailyRoutineCode
	order        RoutineActOrder
	actVersion   act.ActVersionCode
	stage        RoutineActStage
	repsOrMeters RepsOrMeters
	ratioOrSecs  RatioOrSecs
}

func (ra RoutineAct) DailyRoutine() DailyRoutineCode {
	return ra.dailyRoutine
}

func (ra RoutineAct) Order() RoutineActOrder {
	return ra.order
}

func (ra *RoutineAct) setDailyRoutine(drcode DailyRoutineCode) {
	ra.dailyRoutine = drcode
}
