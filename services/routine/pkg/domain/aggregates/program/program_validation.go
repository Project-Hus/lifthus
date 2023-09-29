package program

import "routine/pkg/domain"

const (
	TITLE_MIN_LENGTH = domain.PROGRAM_TITLE_MIN_LENGTH
	TITLE_MAX_LENGTH = domain.PROGRAM_TITLE_MAX_LENGTH
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
