package dto

import (
	"time"
)

/* QUERYING */

/* CREATING */
// to create weekly program rec
type CreateWeeklyProgramRecDto struct {
	Author    uint64    `json:"author,omitempty"`
	ProgramID uint64    `json:"program_id,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`

	Comment string `json:"comment,omitempty"`

	WeeklyRoutineRecs []CreateWeeklyRoutineRecDto `json:"weekly_routine_recs,omitempty"`
}
type CreateWeeklyRoutineRecDto struct {
	WeeklyRoutineID uint64                           `json:"weekly_routine_id,omitempty"`
	Week            int                              `json:"week,omitempty"`
	StartDate       time.Time                        `json:"start_date,omitempty"`
	DayRoutineRecs  []CreateWeeklyDailyRoutineRecDto `json:"daily_routine_recs,omitempty"`
}
type CreateWeeklyDailyRoutineRecDto struct {
	DailyRoutineID uint64    `json:"daily_routine_id,omitempty"`
	Date           time.Time `json:"date,omitempty"`

	RoutineActRecs []CreateWeeklyRoutineActRecDto `json:"routine_act_recs,omitempty"`
}
type CreateWeeklyRoutineActRecDto struct {
	RoutineActID uint64 `json:"routine_act_id,omitempty"`

	ActID uint64 `json:"act_id,omitempty"`
	Order int    `json:"order,omitempty"`
	Reps  *int   `json:"reps,omitempty"`
	Lap   *int   `json:"lap,omitempty"`
}
