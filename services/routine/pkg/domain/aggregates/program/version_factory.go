package program

import (
	"routine/pkg/domain"
	"time"
)

func CreateProgramVersion(
	program ProgramCode,
	version ProgramVersionNumber,
	imageSrcs ProgramImageSrcs,
	text ProgramText,
	dailyRoutines DailyRoutines,
) (*ProgramVersion, error) {
	if !dailyRoutines.IsValid() {
		return nil, ErrInvalidDailyRoutines
	}

	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}

	for _, dr := range dailyRoutines {
		dr.setVersion(ProgramVersionCode(code))
	}

	return ProgramVersionFrom(
		ProgramVersionCode(code),
		program,
		version,
		domain.CreatedAt(time.Now()),
		imageSrcs,
		text,
		dailyRoutines,
	)
}

func ProgramVersionFrom(
	code ProgramVersionCode,
	program ProgramCode,
	version ProgramVersionNumber,
	createdAt domain.CreatedAt,
	imageSrcs ProgramImageSrcs,
	text ProgramText,
	dailyRoutines DailyRoutines,
) (*ProgramVersion, error) {
	if !dailyRoutines.IsValid() {
		return nil, ErrInvalidDailyRoutines
	}
	return &ProgramVersion{
		code:          code,
		program:       program,
		version:       version,
		createdAt:     createdAt,
		imageSrcs:     imageSrcs,
		text:          text,
		dailyRoutines: dailyRoutines,
	}, nil
}
