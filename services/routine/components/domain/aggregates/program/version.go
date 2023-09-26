package program

import (
	"routine/components/domain"
)

const (
	PROGRAM_VERSION_DESCRIPTION_MAX_LENGTH = domain.PROGRAM_VERSION_DESCRIPTION_MAX_LENGTH
)

func isProgramVersionDescValid(desc string) bool {
	return len(desc) <= PROGRAM_VERSION_DESCRIPTION_MAX_LENGTH
}

func DailyVersionFrom(
	description string,
) *Version {
	return &Version{
		descriptiion: description,
	}
}

func WeeklyVersionFrom(
	description string,
) *Version {
	return &Version{
		descriptiion: description,
	}
}

type ProgramVersionNumber int

type Version struct {
	version       ProgramVersionNumber
	descriptiion  string
	weeklyRoutine *[]WeeklyRoutine
	dailyRoutine  *[]DailyRoutine
}
