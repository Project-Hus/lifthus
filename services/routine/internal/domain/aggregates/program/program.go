package program

import (
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
)

type ProgramCode domain.Code

type ProgramTitle string

type Program struct {
	code ProgramCode

	programType ProgramType
	title       ProgramTitle
	author      user.UserId

	createdAt domain.CreatedAt

	parent *ParentProgramVersion

	releases ProgramReleases
}

type ParentProgramVersion struct {
	ProgramCode
	ProgramVersionNumber
}

func (p Program) LatestRelease() *ProgramRelease {
	return p.releases[len(p.releases)-1]
}

func (p Program) Code() ProgramCode {
	return p.code
}

func (p Program) ProgramType() ProgramType {
	return p.programType
}

func (p Program) Title() ProgramTitle {
	return p.title
}

func (p Program) Author() user.UserId {
	return p.author
}

func (p Program) CreatedAt() domain.CreatedAt {
	return p.createdAt
}

func (p Program) ParentProgramVersion() *ParentProgramVersion {
	return p.parent
}

func (p Program) Releases() ProgramReleases {
	return p.releases
}
