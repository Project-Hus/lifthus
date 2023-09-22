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

func (m ProgramMetadata) Ids() (*uint64, string) {
	return m.id, m.code
}

func (m ProgramMetadata) Author() user.User {
	return m.author
}

func (m ProgramMetadata) Timestamps() domain.Timestamps {
	return m.timestamps
}
