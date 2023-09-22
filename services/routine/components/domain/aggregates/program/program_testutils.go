package program

func isCreationSuccess(
	md ProgramMetadata,
	derivedFrom *Program,
	contents ProgramContents,
) bool {
	_, err := Create(
		md.Author(),
		derivedFrom,
		contents,
	)
	return err == nil
}

func getOverflownTitle() string {
	title := ""
	for i := 0; i < TITLE_MAX_LENGTH+1; i++ {
		title += "*"
	}
	return title
}

func getOverflownImageSrcs() []string {
	imageSrcs := []string{}
	for i := 0; i < IMAGES_MAX_COUNT+1; i++ {
		imageSrcs = append(imageSrcs, "https://www.google.com")
	}
	return imageSrcs
}

func getOverflownDescription() string {
	description := ""
	for i := 0; i < DESCRIPTION_MAX_LENGTH+1; i++ {
		description += "*"
	}
	return description
}
