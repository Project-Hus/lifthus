package dto

/* QUERYING */

/* CREATING */
// to create weekly program rec
type CreateWeeklyProgramRecDto struct {
	Title       string  `json:"title,omitempty"`
	Author      uint64  `json:"author,omitempty"`
	Image       *string `json:"image,omitempty"`
	Description *string `json:"description,omitempty"`

	Tags []string `json:"tags,omitempty"`

	WeeklyRoutines []CreateWeeklyRoutineDto      `json:"weekly_routines,omitempty"`
	DailyRoutines  []CreateWeeklyDailyRoutineDto `json:"daily_routines,omitempty"`
	RoutineActs    []CreateWeeklyRoutineActDto   `json:"routine_acts,omitempty"`
}
type CreateWeeklyRoutineRecDto struct {
	Week int `json:"week,omitempty"`
}
type CreateWeeklyDailyRoutineRecDto struct {
	Week int `json:"week,omitempty"`
	Day  int `json:"day,omitempty"`
}
type CreateWeeklyRoutineActRecDto struct {
	Week int `json:"week,omitempty"`
	Day  int `json:"day,omitempty"`

	ActID uint64 `json:"act_id,omitempty"`
	Order int    `json:"order,omitempty"`

	WRatio *float64 `json:"w_ratio,omitempty"`
	Reps   *int     `json:"reps,omitempty"`
	Lap    *int     `json:"lap,omitempty"`
	Warmup bool     `json:"warmup,omitempty"`
}
