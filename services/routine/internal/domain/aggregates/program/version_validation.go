package program

import "routine/internal/domain"

const (
	IMAGES_MIN_NUMBER = domain.PROGRAM_IMAGES_MIN_NUMBER
	IMAGES_MAX_NUMBER = domain.PROGRAM_IMAGES_MAX_NUMBER
	TEXT_MIN_LENGTH   = domain.PROGRAM_TEXT_MIN_LENGTH
	TEXT_MAX_LENGTH   = domain.PROGRAM_TEXT_MAX_LENGTH
)

func (pvs ProgramVersions) IsValid() bool {
	if len(pvs) == 0 {
		return false
	}
	vCnt := ProgramVersionNumber(0)
	for _, pv := range pvs {
		if pv.Version() <= vCnt {
			return false
		}
		vCnt = pv.Version()
	}
	return true
}

func (imgs ProgramImageSrcs) IsValid() bool {
	if len(imgs) > IMAGES_MAX_NUMBER || len(imgs) < IMAGES_MIN_NUMBER {
		return false
	}
	return true
}

func (text ProgramText) IsValid() bool {
	if len(text) > TEXT_MAX_LENGTH || len(text) < TEXT_MIN_LENGTH {
		return false
	}
	return true
}
