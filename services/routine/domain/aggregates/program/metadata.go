package domain

import domain "routine/domain/aggregates/user"

type ProgramMetadata struct {
	code   string
	parent *Program
	title  string
	author domain.User
}

func ProgramMetadataFrom(
	code string,
	parent *Program,
	author domain.User,
	title string,
) ProgramMetadata {
	return ProgramMetadata{
		code:   code,
		parent: parent,
		author: author,
		title:  title,
	}
}

func (pm ProgramMetadata) Code() string {
	return pm.code
}

func (pm ProgramMetadata) Parent() *Program {
	return pm.parent
}

func (pm ProgramMetadata) Author() domain.User {
	return pm.author
}

func (pm ProgramMetadata) Title() string {
	return pm.title
}
