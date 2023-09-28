package program

import (
	"routine/pkg/domain"
	"routine/pkg/domain/aggregates/user"
)

type ProgramCode domain.Code

type ProgramTitle string

type ProgramVersions []*ProgramVersion

type Program struct {
	code ProgramCode

	programType ProgramType
	title       ProgramTitle
	author      user.UserId

	createdAt domain.CreatedAt

	derivedFrom *ProgramVersionCode

	versions ProgramVersions
}
