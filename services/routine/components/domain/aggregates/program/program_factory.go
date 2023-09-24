package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"time"
)

func CreateProgram(
	author user.User,
	programType ProgramType,
	descriptions ProgramDescriptions,
) (*Program, error) {
	if !descriptions.ValidateCreation() {
		return nil, ErrInvalidDescriptions
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	md := MetadataFrom(nil, code, author, domain.TimestampsFrom(time.Now(), nil))
	newProgram := ProgramFrom(md, programType, descriptions)
	return &newProgram, nil
}

func ProgramFrom(
	md ProgramMetadata,
	programType ProgramType,
	descriptions ProgramDescriptions,
) Program {
	return Program{
		metadata:     md,
		programType:  programType,
		descriptions: descriptions,
	}
}
