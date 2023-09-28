package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

type ProgramId domain.Id
type ProgramCode domain.Code

type ProgramTitle string

type Program struct {
	id   *ProgramId
	code ProgramCode

	author      user.UserId
	programType ProgramType
	createdAt   domain.CreatedAt
	updatedAt   *domain.UpdatedAt

	derivedFrom *ProgramId
	deriving    []ProgramId

	title     ProgramTitle
	imageSrcs ProgramImageSrcs
	text      ProgramText

	versions []*Version
}

func (p *Program) Update(updater *user.User, updates ProgramUpdates) (*Program, error) {
	if p.author != updater.Id() {
		return nil, domain.ErrUnauthorized
	}
	switch {
	case updates.Title != nil:
		p.title = *updates.Title
	case updates.ImageSrcs != nil:
		p.imageSrcs = *updates.ImageSrcs
	case updates.Text != nil:
		p.text = *updates.Text
	}
	return p, nil
}

func (p *Program) GetUpdateTargets() ProgramUpdateTargets {
	return ProgramUpdateTargets{
		Title:     p.title,
		ImageSrcs: p.imageSrcs,
		Text:      p.text,
	}
}

func (p *Program) Delete(deleter *user.User) (*Program, error) {
	if p.author != deleter.Id() {
		return nil, domain.ErrUnauthorized
	}
	if len(p.deriving) > 0 {
		return nil, ErrExistingDerivingProgram
	}
	return p, nil
}

func (p Program) IsPersisted() bool {
	return p.id != nil
}

func (p Program) Id() ProgramId {
	return *p.id
}

func (p Program) Code() ProgramCode {
	return p.code
}

func (p Program) Metadata() programMetadata {
	return programMetadata{
		Title:       p.title,
		Author:      p.author,
		ProgramType: p.programType,
		CreatedAt:   p.createdAt,
		UpdatedAt:   p.updatedAt,
	}
}

func (p Program) Derivations() programDerivations {
	return programDerivations{
		DerivedFrom: p.derivedFrom,
		Deriving:    p.deriving,
	}
}

func (p Program) Description() programDescription {
	return programDescription{
		ImageSrcs: p.imageSrcs,
		Text:      p.text,
	}
}

func (p Program) Versions() []*Version {
	return p.versions
}
