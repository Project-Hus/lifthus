package program

func DescriptionsFrom(
	title string,
	imageSrcs []string,
	description string,
) ProgramDescriptions {
	return ProgramDescriptions{
		title:       title,
		imageSrcs:   imageSrcs,
		description: description,
	}
}

type ProgramDescriptions struct {
	title       string
	imageSrcs   []string
	description string
}

func (pd ProgramDescriptions) ValidateCreation() bool {
	return isDescTitleValid(pd.title) &&
		isDescImagesValid(pd.imageSrcs) &&
		isDescDescriptionValid(pd.description)
}

func isDescTitleValid(title string) bool {
	return len(title) <= TITLE_MAX_LENGTH
}

func isDescImagesValid(imageSrcs []string) bool {
	return len(imageSrcs) <= IMAGES_MAX_COUNT
}

func isDescDescriptionValid(description string) bool {
	return len(description) <= DESCRIPTION_MAX_LENGTH
}
