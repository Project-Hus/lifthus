package program

import "routine/components/domain"

const (
	PROGRAM_VERSION_DESCRIPTION_MAX_LENGTH = domain.PROGRAM_VERSION_DESCRIPTION_MAX_LENGTH
)

func IsVersionSequenceValid[T Version](versions []T) bool {
	for idx, v := range versions {
		correctVersion := VersionNumber(idx + 1)
		if v.VersionNumber() != correctVersion {
			return false
		}
	}
	return true
}

func IsVersionDescValid(desc string) bool {
	return len(desc) <= PROGRAM_VERSION_DESCRIPTION_MAX_LENGTH
}
