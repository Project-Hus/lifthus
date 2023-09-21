package program

import (
	"routine/components/domain/aggregates/user"
	"time"
)

// CreateProgram creates a new Program entity if the form is valid, before it is persisted.
func CreateProgram() *Program {
	return &Program{}
}

// ProgramFrom reconstitutes a existing Program entity usually from its persisted state.
func ProgramFrom(id uint64) Program {
	return Program{
		id: &id,
	}
}

// Program is the domain model entity that represents a training program.
type Program struct {
	id     *uint64
	slug   string
	code   string
	parent *Program
	title  string
	author user.User

	programType ProgramType
	iteration   int

	imageSrcs   []string
	description string

	createdAt time.Time
	updatedAt time.Time
}

func (p Program) Id() *uint64 {
	return p.id
}

func (p Program) Info() *ProgramInfo {
	return &ProgramInfo{
		Id:          p.id,
		Slug:        p.slug,
		Code:        p.code,
		ParentId:    p.parent.id,
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

type ProgramInfo struct {
	Id          *uint64
	Slug        string
	Code        string
	ParentId    *uint64
	Title       string
	Author      user.User
	ProgramType ProgramType
	Iteration   int
	ImageSrcs   []string
	Description string
	createdAt   time.Time
	updatedAt   time.Time
}
