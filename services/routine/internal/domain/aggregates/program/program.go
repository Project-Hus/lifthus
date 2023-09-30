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
	curV := p.LatestVersion().Version()
	newV := curV + 1
	newPv, err := CreateProgramVersion(
		p.Code(),
		newV,
		imageSrcs,
		text,
		dailyRoutines,
	)
	if err != nil {
		return nil, err
	}
	p.versions = append(p.versions, newPv)
	return p, nil
}

func (p Program) LatestVersion() *ProgramVersion {
	return p.versions[len(p.versions)-1]
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

func (p Program) DerivedFrom() *ProgramVersionCode {
	return p.derivedFrom
}

func (p Program) Versions() ProgramVersions {
	return p.versions
}
