package program

import (
	"routine/components/domain/aggregates/user"
	"routine/shared/constraints"
	"time"
)

const (
	TITLE_MAX_LENGTH       = constraints.PROGRAM_TITLE_MAX_LENGTH
	IMAGES_MAX_COUNT       = constraints.PROGRAM_IMAGES_MAX_COUNT
	DESCRIPTION_MAX_LENGTH = constraints.PROGRAM_DESCRIPTION_MAX_LENGTH
)

// Program is the domain model entity that represents a training program.
type Program struct {
	id     *uint64
	slug   string
	code   string
	parent *uint64
	title  string
	author user.User

	programType ProgramType
	iteration   int

	imageSrcs   []string
	description string

	createdAt *time.Time
	updatedAt *time.Time
}

func (p Program) Info() *ProgramInfo {
	return &ProgramInfo{
		Id:          p.id,
		Slug:        p.slug,
		Code:        p.code,
		Parent:      p.parent,
		Title:       p.title,
		Author:      p.author,
		ProgramType: p.programType,
		Iteration:   p.iteration,
		ImageSrcs:   p.imageSrcs,
		Description: p.description,
		createdAt:   p.createdAt,
		updatedAt:   p.updatedAt,
	}
}
