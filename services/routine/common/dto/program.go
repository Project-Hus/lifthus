package dto

/* QUERYING */
// to query weekly program
type QueryProgramDto struct {
	ID uint64 `json:"id,omitempty"`

	Title       string   `json:"title,omitempty"`
	Type        string   `json:"type,omitempty"`
	Author      uint64   `json:"author,omitempty"`
	Image       *string  `json:"image,omitempty"`
	Description *string  `json:"description,omitempty"`
	Tags        []uint64 `json:"tags,omitempty"`

	WeeklyRoutines []QueryWeeklyRoutineDto `json:"weekly_routines,omitempty"`
}

type QueryWeeklyRoutineDto struct {
	Week          int                    `json:"week,omitempty"`
	DailyRoutines []QueryDailyRoutineDto `json:"daily_routines,omitempty"`
}
type QueryDailyRoutineDto struct {
	Day         int                  `json:"day,omitempty"`
	RoutineActs []QueryRoutineActDto `json:"routine_acts,omitempty"`
}
type QueryRoutineActDto struct {
	ActID  uint64   `json:"act_id,omitempty"`
	Order  int      `json:"order,omitempty"`
	WRatio *float64 `json:"w_ratio,omitempty"`
	Reps   *int     `json:"reps,omitempty"`
	Lap    *int     `json:"lap,omitempty"`
	Warmup bool     `json:"warmup,omitempty"`
}

// to query daily program

/* CREATING */
// to create weekly program
type CreateWeeklyProgramDto struct {
	Title       string  `json:"title,omitempty"`
	Author      uint64  `json:"author,omitempty"`
	Image       *string `json:"image,omitempty"`
	Description *string `json:"description,omitempty"`

	Tags []string `json:"tags,omitempty"`

	WeeklyRoutines []CreateWeeklyRoutineDto      `json:"weekly_routines,omitempty"`
	DailyRoutines  []CreateWeeklyDailyRoutineDto `json:"daily_routines,omitempty"`
	RoutineActs    []CreateWeeklyRoutineActDto   `json:"routine_acts,omitempty"`
}
type CreateWeeklyRoutineDto struct {
	Week int `json:"week,omitempty"`
}
type CreateWeeklyDailyRoutineDto struct {
	Week int `json:"week,omitempty"`
	Day  int `json:"day,omitempty"`
}
type CreateWeeklyRoutineActDto struct {
	Week int `json:"week,omitempty"`
	Day  int `json:"day,omitempty"`

	ActID uint64 `json:"act_id,omitempty"`
	Order int    `json:"order,omitempty"`

	WRatio *float64 `json:"w_ratio,omitempty"`
	Reps   *int     `json:"reps,omitempty"`
	Lap    *int     `json:"lap,omitempty"`
	Warmup bool     `json:"warmup,omitempty"`
}
