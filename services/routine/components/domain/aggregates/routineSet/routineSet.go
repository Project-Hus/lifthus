package routineset

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/program"
	"time"
)

const (
	DESCRIPTION_MAX_LENGTH = domain.ROUTINESET_DESCRIPTION_MAX_LENGTH
)

func CreateRoutineset(
	metadata RoutinesetMetadata,
	description string,
) (*Routineset, error) {
	if !isDescValid(description) {
		return nil, program.ErrTooLongDescription
	}
	md := CreateRoutinesetMetadata(metadata.programId, metadata.version)
	return &Routineset{
		metadata:    *md,
		description: description,
	}, nil
}

func isDescValid(desc string) bool {
	return len(desc) <= DESCRIPTION_MAX_LENGTH
}

func RoutinesetFrom(
	metadata RoutinesetMetadata,
	derivedFrom *uint64,
	description string,
) *Routineset {
	return &Routineset{
		metadata:    metadata,
		derivedFrom: derivedFrom,
		description: description,
	}
}

type Routineset struct {
	metadata    RoutinesetMetadata
	derivedFrom *uint64
	description string
}

func (rs *Routineset) IsPersisted() bool {
	return rs.metadata.id != nil
}

func (rs *Routineset) Id() *uint64 {
	return rs.metadata.id
}

func (rs *Routineset) ProgramId() uint64 {
	return rs.metadata.programId
}

func (rs *Routineset) Version() uint {
	return rs.metadata.version
}

func (rs *Routineset) Description() string {
	return rs.description
}

func (rs *Routineset) CreatedAt() time.Time {
	return rs.metadata.createdAt
}
