package program

import (
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
	"time"
)

func CreateProgram(
	programType ProgramType,
	title ProgramTitle,
	author user.UserId,
	parent *ParentProgramVersion,
	release ProgramRelease,
) (*Program, error) {
	if !title.IsValid() {
		return nil, ErrInvalidProgramTitle
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	return ProgramFrom(
		ProgramCode(code),
		programType,
		title,
		author,
		domain.CreatedAt(time.Now()),
		parent,
		[]*ProgramRelease{&release},
	)
}

func ProgramFrom(
	code ProgramCode,
	programType ProgramType,
	title ProgramTitle,
	author user.UserId,
	createdAt domain.CreatedAt,
	parent *ParentProgramVersion,
	releases ProgramReleases,
) (*Program, error) {
	if !releases.IsValid() {
		return nil, ErrInvalidProgramReleases
	}
	return &Program{
		code:        code,
		programType: programType,
		title:       title,
		author:      author,
		createdAt:   createdAt,
		parent:      parent,
		releases:    releases,
	}, nil
}
