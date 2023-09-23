package programContents

import "routine/components/domain"

func CreateWeeklyRoutineSet(
	description string,
	weeklyRoutines []WeeklyRoutine,
) (*WeeklyRoutineSet, error) {
	routineSet, err := CreateRoutineSet(description)
	if err != nil {
		return nil, err
	}
	return &WeeklyRoutineSet{
		RoutineSet:     *routineSet,
		weeklyRoutines: weeklyRoutines,
	}, nil
}

func WeeklyRoutineSetFrom(
	version uint,
	description string,
	timestamps domain.Timestamps,
	weeklyRoutines []WeeklyRoutine,
) WeeklyRoutineSet {
	return WeeklyRoutineSet{
		RoutineSet:     RoutineSetFrom(version, description, timestamps),
		weeklyRoutines: weeklyRoutines,
	}
}

type WeeklyRoutineSet struct {
	RoutineSet
	weeklyRoutines []WeeklyRoutine
}
