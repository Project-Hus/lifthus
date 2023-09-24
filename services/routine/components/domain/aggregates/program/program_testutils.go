package program

import "routine/components/domain/aggregates/user"

func isCreationSuccess(
	author user.User,
	pt ProgramType,
	ds ProgramDescriptions,
) bool {
	_, err := CreateProgram(author, pt, ds)
	return err == nil
}

func getNormalTitle() string {
	return "Normal Title"
}

func getOverflownTitle() string {
	title := ""
	for i := 0; i < TITLE_MAX_LENGTH+1; i++ {
		title += "*"
	}
	return title
}

func getNormalImageSrcs() []string {
	return []string{"https://www.google.com"}
}

func getOverflownImageSrcs() []string {
	imageSrcs := []string{}
	for i := 0; i < IMAGES_MAX_COUNT+1; i++ {
		imageSrcs = append(imageSrcs, "https://www.google.com")
	}
	return imageSrcs
}

func getNormalDescription() string {
	return "What doens't kill you makes you stronger."
}

func getOverflownDescription() string {
	description := ""
	for i := 0; i < DESCRIPTION_MAX_LENGTH+1; i++ {
		description += "*"
	}
	return description
}
