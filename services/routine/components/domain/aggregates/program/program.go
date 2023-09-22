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

// Program is the domain model entity that represents a training program.
type Program struct {
	metadata    ProgramMetadata
	derivedFrom *Program
	contents    ProgramContents
}

func (p *Program) Metadata() ProgramMetadata {
	return p.metadata
}

func (p *Program) DerivedFrom() *Program {
	return p.derivedFrom
}

func (p *Program) Contents() ProgramContents {
	return p.contents
}

func (p *Program) Author() user.User {
	return p.metadata.author
}

func (p *Program) IsAuthor(u user.User) bool {
	return p.Author().Id() == u.Id()
}

type ProgramContentsUpdates struct {
	Title       *string
	Description *string
}

func (p *Program) UpdateContents(updater user.User, updates ProgramContentsUpdates) (*Program, error) {
	prev := p.Contents()
	updated := ContentsFrom(prev.programType, prev.iteration, prev.title, prev.imageSrcs, prev.description)
	if !p.IsAuthor(updater) {
		return nil, domain.ErrUnauthorized
	}
	p.contents = updated
	return p, nil
}

func (p *Program) Delete(deleter user.User) (*Program, error) {
	if !p.IsAuthor(deleter) {
		return nil, domain.ErrUnauthorized
	}
	return p, nil
}
