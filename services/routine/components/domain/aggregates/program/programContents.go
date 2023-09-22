package program

func ContentsFrom(
	programType ProgramType,
	iteration int,
	title string,
	imageSrcs []string,
	description string,
) ProgramContents {
	return ProgramContents{
		programType: programType,
		iteration:   iteration,
		title:       title,
		imageSrcs:   imageSrcs,
		description: description,
	}
}

type ProgramContents struct {
	programType ProgramType
	iteration   int
	title       string
	imageSrcs   []string
	description string
}

func (pc ProgramContents) ProgramType() (pt ProgramType, iteration int) {
	return pc.programType, pc.iteration
}

func (pc ProgramContents) ImageSrcs() []string {
	return pc.imageSrcs
}

func (pc ProgramContents) TitleAndDesc() (title, description string) {
	return pc.title, pc.description
}

func (pc ProgramContents) IsValid() bool {
	return isContentsTitleValid(pc.title) &&
		isContentsImageSrcsValid(pc.imageSrcs) &&
		isContentsDescriptionValid(pc.description)
}

func isContentsTitleValid(title string) bool {
	return len(title) <= TITLE_MAX_LENGTH
}

func isContentsImageSrcsValid(imageSrcs []string) bool {
	return len(imageSrcs) <= IMAGES_MAX_COUNT
}

func isContentsDescriptionValid(description string) bool {
	return len(description) <= DESCRIPTION_MAX_LENGTH
}
