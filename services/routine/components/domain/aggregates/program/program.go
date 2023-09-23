package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

const (
	TITLE_MAX_LENGTH       = domain.PROGRAM_TITLE_MAX_LENGTH
	IMAGES_MAX_COUNT       = domain.PROGRAM_IMAGES_MAX_COUNT
	DESCRIPTION_MAX_LENGTH = domain.PROGRAM_DESCRIPTION_MAX_LENGTH
)

type Program struct {
	metadata     ProgramMetadata
	derivedFrom  *ProgramDerivedFrom
	programType  ProgramType
	descriptions ProgramDescriptions
}

func (p Program) IdAndCode() (*uint64, string) {
	return p.metadata.id, p.metadata.code
}

func (p Program) Author() user.User {
	return p.metadata.author
}

func (p Program) Timestamps() domain.Timestamps {
	return p.metadata.timestamps
}

func (p Program) DerivedFrom() (program *Program, version uint) {
	if p.derivedFrom == nil {
		return nil, 0
	}
	return p.derivedFrom.program, p.derivedFrom.version
}

func (p Program) Type() (ptype string, iteration uint) {
	return p.programType.ptype, p.programType.iteration
}

func (p Program) ImageSrcs() []string {
	return p.descriptions.imageSrcs
}

func (p Program) TitleAndDescription() (title string, description string) {
	return p.descriptions.title, p.descriptions.description
}
