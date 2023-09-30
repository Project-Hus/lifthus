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

	derivedFrom *ProgramVersionCode

	versions ProgramVersions
}

func (p *Program) Upgrade(author user.User, imageSrcs ProgramImageSrcs, text ProgramText, dailyRoutines DailyRoutines) (*Program, error) {
	return nil, nil
}
