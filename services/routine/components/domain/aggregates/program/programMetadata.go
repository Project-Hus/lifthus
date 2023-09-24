package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

func MetadataFrom(id *uint64, code string, author user.User, timestamps domain.Timestamps) ProgramMetadata {
	return ProgramMetadata{
		id:         id,
		code:       code,
		author:     author,
		timestamps: timestamps,
	}
}

type ProgramMetadata struct {
	id         *uint64
	code       string
	author     user.User
	timestamps domain.Timestamps
}
