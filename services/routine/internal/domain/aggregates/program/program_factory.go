package program

import (
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
	"time"
)

func CreateWeeklyProgram(
	title ProgramTitle,
	author user.User,
	derivedFrom *ProgramVersionCode,
	imageSrcs ProgramImageSrcs,
	text ProgramText,
	dailyRoutines DailyRoutines,
) (*Program, error) {
	if !title.IsValid() {
		return nil, ErrInvalidProgramTitle
	}
	pcode, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	v1, err := CreateProgramVersion(
		ProgramCode(pcode),
		1,
		imageSrcs,
		text,
		dailyRoutines,
	)
	if err != nil {
		return nil, err
	}
	return WeeklyProgramFrom(
		ProgramCode(pcode),
		title,
		author,
		domain.CreatedAt(time.Now()),
		derivedFrom,
		ProgramVersions{v1},
	)
}

func WeeklyProgramFrom(
	code ProgramCode,
	title ProgramTitle,
	author user.User,
	createdAt domain.CreatedAt,
	derivedFrom *ProgramVersionCode,
	versions ProgramVersions,
) (*Program, error) {
	if !versions.IsValid() {
		return nil, ErrInvalidProgramVersions
	}
	return &Program{
		code:        code,
		programType: WeeklyType,
		title:       title,
		author:      author.Id(),
		createdAt:   createdAt,
		derivedFrom: derivedFrom,
		versions:    versions,
	}, nil
}