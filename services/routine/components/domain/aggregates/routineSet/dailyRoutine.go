package routineSet

import "routine/components/domain/aggregates/program"

type DailyRoutine struct {
	id          uint64
	programType program.ProgramType
	parentId    uint64
	order       uint
	routineActs []RoutineAct
}
