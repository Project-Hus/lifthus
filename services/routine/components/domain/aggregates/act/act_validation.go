package act

func IsActNameValid(name ActName) bool {
	return len(name) <= NAME_MAX_LENGTH && len(name) >= NAME_MIN_LENGTH
}

func IsActImagesValid(images []ActImageSrc) bool {
	return len(images) <= IMAGES_MAX_NUMBER
}

func IsActTextValid(text ActText) bool {
	return len(text) <= TEXT_MAX_LENGTH && len(text) >= TEXT_MIN_LENGTH
}

func IsActUpdatesValid(updates ActUpdates) bool {
	switch {
	case updates.ImageSrcs != nil && !IsActImagesValid(*updates.ImageSrcs):
		fallthrough
	case updates.Text != nil && !IsActTextValid(*updates.Text):
		return false
	}
	return true
}
