package domain

import "routine/domain/aggregates/user"

// CreateProgram creates a new Program entity if the form is valid, before it is persisted.
func CreateProgram(md ProgramMetadata, ct ProgramContents, dsc ProgramDescription) *Program {
	if descInvalid(dsc) {
		return nil
	}
	return &Program{}
}

func descInvalid(dsc ProgramDescription) bool {
	return len(dsc.imageSrcs) != 1 || len(dsc.text) > 5000
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
	code   string
	parent *Program
	title  string
	author user.User

	programType ProgramType
	iteration   int

	imageSrcs   []string
	description string
}

func (p Program) Id() *uint64 {
	return p.id
}
