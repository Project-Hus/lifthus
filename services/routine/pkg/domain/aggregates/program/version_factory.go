package program

import "routine/pkg/domain"

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
