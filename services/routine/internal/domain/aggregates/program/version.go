package program

import "routine/internal/domain"

type ProgramVersionCode domain.Code
type ProgramVersionNumber uint
type ProgramImageSrcs []string
type ProgramText string

type ProgramVersions []*ProgramVersion

type ProgramVersion struct {
	code      ProgramVersionCode
	program   ProgramCode
	version   ProgramVersionNumber
	createdAt domain.CreatedAt

	imageSrcs ProgramImageSrcs
	text      ProgramText

	dailyRoutines DailyRoutines
}

func (pv ProgramVersion) Code() ProgramVersionCode {
	return pv.code
}

func (pv ProgramVersion) Program() ProgramCode {
	return pv.program
}

func (pv ProgramVersion) Version() ProgramVersionNumber {
	return pv.version
}

func (pv ProgramVersion) CreatedAt() domain.CreatedAt {
	return pv.createdAt
}

func (pv ProgramVersion) ImageSrcs() ProgramImageSrcs {
	return pv.imageSrcs
}

func (pv ProgramVersion) Text() ProgramText {
	return pv.text
}

func (pv ProgramVersion) DailyRoutines() DailyRoutines {
	return pv.dailyRoutines
}
