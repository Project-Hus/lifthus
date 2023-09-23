package programContents

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/program"
	"time"
)

func CreateRoutineSet(
	description string,
) (*RoutineSet, error) {
	if len(description) > domain.ROUTINESET_DESCRIPTION_MAX_LENGTH {
		return nil, program.ErrTooLongDescription
	}
	return &RoutineSet{
		description: description,
		timestamps:  domain.TimestampsFrom(time.Now(), nil),
	}, nil
}

func RoutineSetFrom(
	version uint,
	description string,
	timestamps domain.Timestamps,
) RoutineSet {
	return RoutineSet{
		version:     version,
		description: description,
		timestamps:  timestamps,
	}
}

type RoutineSet struct {
	version     uint
	description string
	timestamps  domain.Timestamps
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
