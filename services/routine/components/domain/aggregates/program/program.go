package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

type ProgramCode domain.Code

type ProgramTitle string

type Program struct {
	code ProgramCode

	programType ProgramType
	title       ProgramTitle
	author      user.UserId

	createdAt domain.CreatedAt

	derivedFrom *ProgramVersionCode

	versions []*ProgramVersion
}
