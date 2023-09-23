package programContents

import "routine/components/domain/aggregates/program"

type WeeklyRoutine struct {
	id           uint64
	program      program.Program
	order        uint
	dailyroutine [7]DailyRoutine
}
