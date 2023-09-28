package program

import "routine/components/domain"

type ProgramVersionCode domain.Code
type ProgramVersionNumber uint
type ProgramImageSrcs []string
type ProgramText string

type DailyRoutines []*DailyRoutine

type ProgramVersion struct {
	code      ProgramVersionCode
	program   ProgramCode
	version   ProgramVersionNumber
	createdAt domain.CreatedAt

	imageSrcs ProgramImageSrcs
	text      ProgramText

	dailyRoutines DailyRoutines
}
