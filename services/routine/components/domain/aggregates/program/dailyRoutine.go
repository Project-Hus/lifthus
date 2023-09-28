package program

import "routine/components/domain"

type DailyRoutineCode domain.Code

type DailyRoutineDay uint
type DailyRoutineDescription string
type DailyRoutine struct {
	day         DailyRoutineDay
	description DailyRoutineDescription
}
