package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"time"
)

func Create(
	author user.User,
	parent *Program,
	contents ProgramContents,
) (*Program, error) {
	code, err := domain.RandomHex(domain.CODE_LENGTH)
	if err != nil {
		return nil, err
	}

	newMetadata := MetadataFrom(nil, code, author, domain.TimestampsFrom(time.Now(), nil))

	if !contents.IsValid() {
		return nil, ErrInvalidContents
	}

	return &Program{
		metadata:    newMetadata,
		derivedFrom: parent,
		contents:    contents,
	}, nil
}

func From(
	md ProgramMetadata,
	derivedFrom *Program,
	contents ProgramContents,
) *Program {
	return &Program{
		metadata:    md,
		derivedFrom: derivedFrom,
		contents:    contents,
	}
}
