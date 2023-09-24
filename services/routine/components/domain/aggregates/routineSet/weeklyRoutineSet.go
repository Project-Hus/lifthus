package routineSet

import "routine/components/domain"

func CreateWeeklyRoutineSet(
	programId uint64,
	description string,
	weeklyRoutines []WeeklyRoutine,
) (*WeeklyRoutineSet, error) {
	routineSet, err := CreateRoutineSet(programId, description)
	if err != nil {
		return nil, err
	}
	return &WeeklyRoutineSet{
		RoutineSet:     *routineSet,
		weeklyRoutines: weeklyRoutines,
	}, nil
}

func WeeklyRoutineSetFrom(
	id uint64,
	programId uint64,
	version uint,
	description string,
	timestamps domain.Timestamps,
	weeklyRoutines []WeeklyRoutine,
) *WeeklyRoutineSet {
	return &WeeklyRoutineSet{
		RoutineSet:     *RoutineSetFrom(id, programId, version, description, timestamps),
		weeklyRoutines: weeklyRoutines,
	}
}

type WeeklyRoutineSet struct {
	RoutineSet
	weeklyRoutines []WeeklyRoutine
}
