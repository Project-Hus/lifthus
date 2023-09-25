package programset

import (
	"routine/components/domain"
	"time"
)

const (
	DESCRIPTION_MAX_LENGTH = domain.ROUTINESET_DESCRIPTION_MAX_LENGTH
)

func CreateEdition(
	metadata EditionMetadata,
	description string,
) (*Edition, error) {
	if !isDescValid(description) {
		return nil, nil
	}
	md := CreateEditionMetadata(metadata.programId, metadata.version)
	return &Edition{
		metadata:    *md,
		description: description,
	}, nil
}

func isDescValid(desc string) bool {
	return len(desc) <= DESCRIPTION_MAX_LENGTH
}

func EditionFrom(
	metadata EditionMetadata,
	derivedFrom *uint64,
	description string,
) *Edition {
	return &Edition{
		metadata:    metadata,
		derivedFrom: derivedFrom,
		description: description,
	}
}

type Edition struct {
	metadata    EditionMetadata
	derivedFrom *uint64
	description string
}

func (rs *Edition) IsPersisted() bool {
	return rs.metadata.id != nil
}

func (rs *Edition) Id() *uint64 {
	return rs.metadata.id
}

func (rs *Edition) ProgramId() uint64 {
	return rs.metadata.programId
}

func (rs *Edition) Version() uint {
	return rs.metadata.version
}

func (rs *Edition) Description() string {
	return rs.description
}

func (rs *Edition) CreatedAt() time.Time {
	return rs.metadata.createdAt
}
