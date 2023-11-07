package program

import "routine/internal/domain"

type ProgramVersionNumber int
type ProgramImageSrcs []string
type ProgramText string

type ProgramReleases []*ProgramRelease

type ProgramRelease struct {
	version   ProgramVersionNumber
	createdAt domain.CreatedAt
	imageSrcs ProgramImageSrcs
	text      ProgramText

	routines Routines
}

func (pv ProgramRelease) Version() ProgramVersionNumber {
	return pv.version
}

func (pv ProgramRelease) CreatedAt() domain.CreatedAt {
	return pv.createdAt
}

func (pv ProgramRelease) ImageSrcs() ProgramImageSrcs {
	return pv.imageSrcs
}

func (pv ProgramRelease) Text() ProgramText {
	return pv.text
}

func (pv ProgramRelease) Routines() Routines {
	return pv.routines
}
