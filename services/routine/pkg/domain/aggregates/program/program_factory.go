package program

import (
	"routine/pkg/domain"
	"routine/pkg/domain/aggregates/user"
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
	pcode, vcode, err := getProgramCodeAndVersionCode()
	if err != nil {
		return nil, err
	}
	v1, err := ProgramVersionFrom(
		ProgramVersionCode(vcode),
		ProgramCode(pcode),
		ProgramVersionNumber(1),
		domain.CreatedAt(time.Now()),
		imageSrcs,
		text,
		dailyRoutines,
	)
	if err != nil {
		return nil, err
	}
	versions := ProgramVersions{v1}
	return WeeklyProgramFrom(
		ProgramCode(pcode),
		title,
		author,
		domain.CreatedAt(time.Now()),
		derivedFrom,
		versions,
	)
}

func getProgramCodeAndVersionCode() (ProgramCode, ProgramVersionCode, error) {
	pCode, err := domain.RandomHexCode()
	if err != nil {
		return "", "", err
	}
	vCode, err := domain.RandomHexCode()
	if err != nil {
		return "", "", err
	}
	return ProgramCode(pCode), ProgramVersionCode(vCode), nil
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
