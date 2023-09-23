package routineSet

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/program"
	"time"
)

const (
	DESCRIPTION_MAX_LENGTH = domain.ROUTINESET_DESCRIPTION_MAX_LENGTH
)

func CreateRoutineSet(
	programId uint64,
	description string,
) (*RoutineSet, error) {
	if !isDescValid(description) {
		return nil, program.ErrTooLongDescription
	}
	return &RoutineSet{
		programId:   programId,
		description: description,
		timestamps:  domain.TimestampsFrom(time.Now(), nil),
	}, nil
}

func isDescValid(desc string) bool {
	return len(desc) <= DESCRIPTION_MAX_LENGTH
}

func RoutineSetFrom(
	id uint64,
	programId uint64,
	version uint,
	description string,
	timestamps domain.Timestamps,
) *RoutineSet {
	return &RoutineSet{
		id:          &id,
		programId:   programId,
		version:     version,
		description: description,
		timestamps:  timestamps,
	}
}

type RoutineSet struct {
	id          *uint64
	programId   uint64
	version     uint
	description string
	timestamps  domain.Timestamps
}

func (rs *RoutineSet) IsPersisted() bool {
	return rs.id != nil
}

func (rs *RoutineSet) Id() *uint64 {
	return rs.id
}

func (rs *RoutineSet) ProgramId() uint64 {
	return rs.programId
}

func (rs *RoutineSet) Version() uint {
	return rs.version
}

func (rs *RoutineSet) Description() string {
	return rs.description
}

func (rs *RoutineSet) CreatedAt() time.Time {
	return rs.timestamps.CreatedAt()
}
