package program

import "routine/components/domain"

type VersionNumber int
type VersionDescription string
type RoutineDescription string

type Version struct {
	version     VersionNumber
	description VersionDescription
	createdAt   domain.CreatedAt
	routineDescription map[RoutineDescription
	dailyRoutines  *[]*DailyRoutine
}

func (v Version) VersionNumber() VersionNumber {
	return v.version
}

func (v Version) Description() VersionDescription {
	return v.description
}

func (v Version) CreatedAt() domain.CreatedAt {
	return v.createdAt
}