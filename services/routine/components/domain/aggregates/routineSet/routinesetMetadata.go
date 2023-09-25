package routineset

import "time"

func CreateRoutinesetMetadata(
	programId uint64,
	version uint,
) *RoutinesetMetadata {
	return &RoutinesetMetadata{
		programId: programId,
		version:   version,
		createdAt: time.Now(),
	}
}

func RoutinesetMetadataFrom(
	id uint64,
	programId uint64,
	version uint,
	createdAt time.Time,
) *RoutinesetMetadata {
	return &RoutinesetMetadata{
		id:        &id,
		programId: programId,
		version:   version,
		createdAt: createdAt,
	}
}

type RoutinesetMetadata struct {
	id        *uint64
	programId uint64
	version   uint
	createdAt time.Time
}
