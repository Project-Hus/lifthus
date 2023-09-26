package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

type ProgramId uint64
type ProgramCode string

type ProgramTitle string
type ProgramImageSrcs []string
type ProgramDescription string

const (
	TITLE_MAX_LENGTH       = domain.PROGRAM_TITLE_MAX_LENGTH
	IMAGES_MAX_COUNT       = domain.PROGRAM_IMAGES_MAX_COUNT
	DESCRIPTION_MAX_LENGTH = domain.PROGRAM_DESCRIPTION_MAX_LENGTH
)

func IsTitleValid(title ProgramTitle) bool {
	return len(title) <= TITLE_MAX_LENGTH
}

func IsImageSrcsValid(imageSrcs ProgramImageSrcs) bool {
	return len(imageSrcs) <= IMAGES_MAX_COUNT
}

func IsDescriptionValid(description ProgramDescription) bool {
	return len(description) <= DESCRIPTION_MAX_LENGTH
}

type Program struct {
	id   *ProgramId
	code ProgramCode

	author      user.UserId
	programType ProgramType
	createdAt   domain.CreatedAt
	updatedAt   *domain.UpdatedAt

	derivedFrom *ProgramId
	deriving    []ProgramId

	title       ProgramTitle
	imageSrcs   ProgramImageSrcs
	description ProgramDescription
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
	case updates.Description != nil:
		p.description = *updates.Description
	}
	return p, nil
}

type ProgramUpdates struct {
	Title       *ProgramTitle
	ImageSrcs   *ProgramImageSrcs
	Description *ProgramDescription
}

func (p *Program) GetUpdateTargets() ProgramUpdateTargets {
	return ProgramUpdateTargets{
		Title:       p.title,
		ImageSrcs:   p.imageSrcs,
		Description: p.description,
	}
}

type ProgramUpdateTargets struct {
	Title       ProgramTitle
	ImageSrcs   ProgramImageSrcs
	Description ProgramDescription
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

func (p Program) Info() programInfo {
	return programInfo{
		Title:       p.title,
		ImageSrcs:   p.imageSrcs,
		Description: p.description,
	}
}
