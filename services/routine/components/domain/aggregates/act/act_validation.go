package act

func IsNewActValid(base ActBase, desc ActDescription) bool {
	return IsActNameValid(base.ActName) &&
		IsActImagesValid(desc.ImageSrcs) &&
		IsActTextValid(desc.Text)
}

func IsActNameValid(name ActName) bool {
	return len(name) <= NAME_MAX_LENGTH
}

func IsActImagesValid(images ActImageSrcs) bool {
	return len(images) <= IMAGES_MAX_NUMBER
}

func IsActTextValid(text ActText) bool {
	return len(text) <= TEXT_MAX_LENGTH
}
