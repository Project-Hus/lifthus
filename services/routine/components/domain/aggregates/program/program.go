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
	programType  ProgramType
	descriptions ProgramDescriptions
}

type ProgramUpdates struct {
	Title       *string
	ImageSrcs   *[]string
	Description *string
}

func (p *Program) UpdateProgramType(updates ProgramUpdates) *Program {
	switch {
	case updates.Title != nil:
		p.descriptions.title = *updates.Title
	case updates.ImageSrcs != nil:
		p.descriptions.imageSrcs = *updates.ImageSrcs
	case updates.Description != nil:
		p.descriptions.description = *updates.Description
	}
	return p
}

type ProgramUpdateTargets struct {
	Title       string
	ImageSrcs   []string
	Description string
}

func (p *Program) GetUpdateTargets() ProgramUpdateTargets {
	return ProgramUpdateTargets{
		Title:       p.descriptions.title,
		ImageSrcs:   p.descriptions.imageSrcs,
		Description: p.descriptions.description,
	}
}

func (p *Program) Delete() *Program {
	return p
}

func (p Program) IsPersisted() bool {
	return p.metadata.id != nil
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

func (p Program) Type() ProgramType {
	return p.programType
}

func (p Program) ImageSrcs() []string {
	return p.descriptions.imageSrcs
}

func (p Program) TitleAndDescription() (title string, description string) {
	return p.descriptions.title, p.descriptions.description
}
