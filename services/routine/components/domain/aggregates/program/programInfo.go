package program

func ProgramInfoFrom(
	title ProgramTitle,
	imageSrcs ProgramImageSrcs,
	description ProgramDescription,
) programInfo {
	return programInfo{
		Title:       title,
		ImageSrcs:   imageSrcs,
		Description: description,
	}
}

type programInfo struct {
	Title       ProgramTitle
	ImageSrcs   ProgramImageSrcs
	Description ProgramDescription
}
