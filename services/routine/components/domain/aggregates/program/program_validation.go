package program

import "routine/components/domain"

const (
	TITLE_MAX_LENGTH               = domain.PROGRAM_TITLE_MAX_LENGTH
	IMAGES_MAX_COUNT               = domain.PROGRAM_IMAGES_MAX_COUNT
	PROGRAM_DESCRIPTION_MAX_LENGTH = domain.PROGRAM_DESCRIPTION_MAX_LENGTH
)

func IsProgramInfoValid(info programInfo) bool {
	return IsTitleValid(info.Title) &&
		IsImageSrcsValid(info.ImageSrcs) &&
		IsDescriptionValid(info.Description)
}

func IsTitleValid(title ProgramTitle) bool {
	return len(title) <= TITLE_MAX_LENGTH
}

func IsImageSrcsValid(imageSrcs ProgramImageSrcs) bool {
	return len(imageSrcs) <= IMAGES_MAX_COUNT
}

func IsDescriptionValid(description ProgramDescription) bool {
	return len(description) <= PROGRAM_DESCRIPTION_MAX_LENGTH
}
