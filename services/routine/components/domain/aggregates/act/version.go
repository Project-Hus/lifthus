package act

import "routine/components/domain"

const (
	NAME_MIN_LENGTH   = domain.ACT_NAME_MIN_LENGTH
	NAME_MAX_LENGTH   = domain.ACT_NAME_MAX_LENGTH
	IMAGES_MAX_NUMBER = domain.ACT_IMAGES_MAX_NUMBER
	TEXT_MIN_LENGTH   = domain.ACT_TEXT_MIN_LENGTH
	TEXT_MAX_LENGTH   = domain.ACT_TEXT_MAX_LENGTH
)

type ActVersionNumber uint
type ActImageSrcs []string
type ActText string

type ActVersion struct {
	version         ActVersionNumber
	imageSrcs       ActImageSrcs
	text            ActText
	characteristics ActCharacteristics
	createdAt       domain.CreatedAt
}

func (v ActVersion) Version() ActVersionNumber {
	return v.version
}

func (v ActVersion) ImageSrcs() ActImageSrcs {
	return v.imageSrcs
}

func (v ActVersion) Text() ActText {
	return v.text
}

func (v ActVersion) Characteristics() ActCharacteristics {
	return v.characteristics
}

func (v ActVersion) CreatedAt() domain.CreatedAt {
	return v.createdAt
}
