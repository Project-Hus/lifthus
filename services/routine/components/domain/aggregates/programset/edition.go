package programset

import (
	"routine/components/domain"
)

const (
	DESCRIPTION_MAX_LENGTH = domain.ROUTINESET_DESCRIPTION_MAX_LENGTH
)

func isDescValid(desc string) bool {
	return len(desc) <= DESCRIPTION_MAX_LENGTH
}

func EditionFrom(
	derivedFrom *uint64,
	description string,
) *Edition {
	return &Edition{
		derivedFrom: derivedFrom,
		description: description,
	}
}

type Edition struct {
	derivedFrom *uint64
	description string
}
