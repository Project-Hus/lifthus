package program

import "routine/internal/domain/aggregates/act"

func CreateRoutineActWithoutDailyRoutine(
	order RoutineActOrder,
	actCode act.ActCode,
	stage RoutineActStage,
	repsOrMeters RepsOrMeters,
	ratioOrSecs RatioOrSecs,
) *RoutineAct {
	return &RoutineAct{
		order:        order,
		actCode:      actCode,
		stage:        stage,
		repsOrMeters: repsOrMeters,
		ratioOrSecs:  ratioOrSecs,
	}
}

func RoutineActFrom(
	order RoutineActOrder,
	actCode act.ActCode,
	stage RoutineActStage,
	repsOrMeters RepsOrMeters,
	ratioOrSecs RatioOrSecs,
) *RoutineAct {
	return &RoutineAct{
		order:        order,
		actCode:      actCode,
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
	order        RoutineActOrder
	actCode      act.ActCode
	stage        RoutineActStage
	repsOrMeters RepsOrMeters
	ratioOrSecs  RatioOrSecs
}

func (ra RoutineAct) Order() RoutineActOrder {
	return ra.order
}

func (ra RoutineAct) Act() act.ActCode {
	return ra.actCode
}

func (ra RoutineAct) Stage() RoutineActStage {
	return ra.stage
}

func (ra RoutineAct) RepsOrMeters() RepsOrMeters {
	return ra.repsOrMeters
}

func (ra RoutineAct) RatioOrSecs() RatioOrSecs {
	return ra.ratioOrSecs
}
