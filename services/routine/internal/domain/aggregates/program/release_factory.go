package program

import (
	"routine/internal/domain"
	"time"
)

func CreateProgramRelease(
	version ProgramVersionNumber,
	imageSrcs ProgramImageSrcs,
	text ProgramText,
	routines Routines,
) (*ProgramRelease, error) {
	if !imageSrcs.IsValid() {
		return nil, ErrInvalidProgramImageSrcs
	}
	if !text.IsValid() {
		return nil, ErrInvalidProgramText
	}

	return ProgramReleaseFrom(
		version,
		domain.CreatedAt(time.Now()),
		imageSrcs,
		text,
		routines,
	)
}

func ProgramReleaseFrom(
	version ProgramVersionNumber,
	createdAt domain.CreatedAt,
	imageSrcs ProgramImageSrcs,
	text ProgramText,
	routines Routines,
) (*ProgramRelease, error) {
	if !routines.IsValid() {
		return nil, ErrInvalidRoutines
	}
	return &ProgramRelease{
		version:   version,
		createdAt: createdAt,
		imageSrcs: imageSrcs,
		text:      text,
		routines:  routines,
	}, nil
}
