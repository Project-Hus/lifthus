package act

import "routine/internal/domain"

const (
	NAME_MIN_LENGTH   = domain.ACT_NAME_MIN_LENGTH
	NAME_MAX_LENGTH   = domain.ACT_NAME_MAX_LENGTH
	IMAGES_MAX_NUMBER = domain.ACT_IMAGES_MAX_NUMBER
	TEXT_MIN_LENGTH   = domain.ACT_TEXT_MIN_LENGTH
	TEXT_MAX_LENGTH   = domain.ACT_TEXT_MAX_LENGTH
)

func (name ActName) IsValid() bool {
	return len(name) <= NAME_MAX_LENGTH && len(name) >= NAME_MIN_LENGTH
}

func (images ActImageSrcs) IsValid() bool {
	return len(images) <= IMAGES_MAX_NUMBER
}

func (text ActText) IsValid() bool {
	return len(text) <= TEXT_MAX_LENGTH && len(text) >= TEXT_MIN_LENGTH
}

func (ut ActUpdateTargets) IsValid() bool {
	if ut.ImageSrcs != nil && !ut.ImageSrcs.IsValid() {
		return false
	}
	if ut.Text != nil && !ut.Text.IsValid() {
		return false
	}
	return true
}
