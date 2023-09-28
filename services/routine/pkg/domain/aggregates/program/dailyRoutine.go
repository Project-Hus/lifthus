package program

import "routine/pkg/domain"

type DailyRoutineCode domain.Code
type DailyRoutineDay uint

type DailyRoutine struct {
	code DailyRoutineCode
	day  DailyRoutineDay

	routineActs []*RoutineAct
}
