package domain

type ProgramDescription struct {
	imageSrcs []string
	text      string
}

func ProgramDescriptionFrom(
	imageSrcs []string,
	text string,
) ProgramDescription {
	return ProgramDescription{
		imageSrcs: imageSrcs,
		text:      text,
	}
}

func (pd ProgramDescription) ImageSrcs() []string {
	return pd.imageSrcs
}

func (pd ProgramDescription) Text() string {
	return pd.text
}
