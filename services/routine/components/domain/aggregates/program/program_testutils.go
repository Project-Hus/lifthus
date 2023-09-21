package program

func IsCreationSuccess(info *NewProgramInfo) bool {
	_, err := Create(info)
	return err == nil
}

func GetOverflownTitle() string {
	title := ""
	for i := 0; i < TITLE_MAX_LENGTH+1; i++ {
		title += "*"
	}
	return title
}

func GetOverflownImageSrcs() []string {
	imageSrcs := []string{}
	for i := 0; i < IMAGES_MAX_COUNT+1; i++ {
		imageSrcs = append(imageSrcs, "https://www.google.com")
	}
	return imageSrcs
}

func GetOverflownDescription() string {
	description := ""
	for i := 0; i < DESCRIPTION_MAX_LENGTH+1; i++ {
		description += "*"
	}
	return description
}
