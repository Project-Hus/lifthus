package programset

import "time"

func CreateEditionMetadata(
	programId uint64,
	version uint,
) *EditionMetadata {
	return &EditionMetadata{
		programId: programId,
		version:   version,
		createdAt: time.Now(),
	}
}

func EditionMetadataFrom(
	id uint64,
	programId uint64,
	version uint,
	createdAt time.Time,
) *EditionMetadata {
	return &EditionMetadata{
		id:        &id,
		programId: programId,
		version:   version,
		createdAt: createdAt,
	}
}

type EditionMetadata struct {
	id        *uint64
	programId uint64
	version   uint
	createdAt time.Time
}
