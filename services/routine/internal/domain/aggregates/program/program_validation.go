package program

import "routine/internal/domain"

const (
	TITLE_MIN_LENGTH = domain.PROGRAM_TITLE_MIN_LENGTH
	TITLE_MAX_LENGTH = domain.PROGRAM_TITLE_MAX_LENGTH
)

func (pt ProgramTitle) IsValid() bool {
	if len(pt) > TITLE_MAX_LENGTH || len(pt) < TITLE_MIN_LENGTH {
		return false
	}
	return true
}
