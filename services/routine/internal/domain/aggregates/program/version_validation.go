package program

import "routine/internal/domain"

const (
	IMAGES_MIN_NUMBER = domain.PROGRAM_IMAGES_MIN_NUMBER
	IMAGES_MAX_NUMBER = domain.PROGRAM_IMAGES_MAX_NUMBER
	TEXT_MIN_LENGTH   = domain.PROGRAM_TEXT_MIN_LENGTH
	TEXT_MAX_LENGTH   = domain.PROGRAM_TEXT_MAX_LENGTH
)

func (pvs ProgramVersions) IsValid() bool {
	vCnt := ProgramVersionNumber(0)
	for _, pv := range pvs {
		if pv.Version() <= vCnt {
			return false
		}
		vCnt = pv.Version()
	}
	return true
}
