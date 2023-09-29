package program

import "routine/pkg/domain"

const (
	IMAGES_MIN_NUMBER = domain.PROGRAM_IMAGES_MIN_NUMBER
	IMAGES_MAX_NUMBER = domain.PROGRAM_IMAGES_MAX_NUMBER
	TEXT_MIN_LENGTH   = domain.PROGRAM_TEXT_MIN_LENGTH
	TEXT_MAX_LENGTH   = domain.PROGRAM_TEXT_MAX_LENGTH
)

func (drs DailyRoutines) IsValid() bool {
	dCnt := DailyRoutineDay(0)
	for _, dr := range drs {
		if dr.Day() <= dCnt {
			return false
		}
		dCnt = dr.day
	}
	return true
}
