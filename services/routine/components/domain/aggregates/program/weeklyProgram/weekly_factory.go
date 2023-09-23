package weeklyProgram

import (
	"routine/components/domain/aggregates/program"
	"routine/components/domain/aggregates/user"
)

func CreateWeeklyProgram(
	author user.User,
	derivedFrom *program.ProgramDerivedFrom,
	programType program.ProgramType,
	descriptions program.ProgramDescriptions,
) (*WeeklyProgram, error) {
	newProgram, err := program.CreateProgram(
		author, derivedFrom, programType, descriptions,
	)
	if err != nil {
		return nil, err
	}
	newWeeklyProgram := &WeeklyProgram{
		Program: *newProgram,
	}
	return newWeeklyProgram, nil
}

func WeeklyProgramFrom() WeeklyProgram {
	newProgram := program.Program{}
	newWeeklyProgram := WeeklyProgram{
		Program: newProgram,
	}
	return newWeeklyProgram
}
