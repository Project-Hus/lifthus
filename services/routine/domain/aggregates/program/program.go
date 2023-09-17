package domain

import domain "routine/domain/shared"

// CreateProgram creates a new Program entity if the form is valid, before it is persisted.
func CreateProgram(md ProgramMetadata, ct ProgramContents, dsc ProgramDescription) *Program {
	if descInvalid(dsc) {
		return nil
	}
	return &Program{
		metadata:    md,
		contents:    ct,
		description: dsc,
	}
}

func descInvalid(dsc ProgramDescription) bool {
	return len(dsc.imageSrcs) != 1 || len(dsc.text) > 5000
}

// ProgramFrom reconstitutes a existing Program entity usually from its persisted state.
func ProgramFrom(id uint64, md ProgramMetadata, ct ProgramContents, dsc ProgramDescription, ts domain.Timestamps) Program {
	return Program{
		id:          &id,
		metadata:    md,
		contents:    ct,
		description: dsc,
		timestamps:  &ts,
	}
}

// Program is the domain model entity that represents a training program.
type Program struct {
	id *uint64

	metadata    ProgramMetadata
	contents    ProgramContents
	description ProgramDescription

	timestamps *domain.Timestamps
}

func (u Program) Id() *uint64 {
	return u.id
}

func (u Program) Metadata() ProgramMetadata {
	return u.metadata
}
