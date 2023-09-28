package act

import "routine/components/domain"

func getValidActDescription() ActDescription {
	return ActDescriptionFrom(getValidActImages(), getValidActText(), getValidCharacteristics())
}

func getValidActName() ActName {
	return "High Bar Squat"
}

func getTooLongActName() ActName {
	ln := ""
	for i := 0; i < domain.ACT_NAME_MAX_LENGTH+1; i++ {
		ln += "a"
	}
	return ActName(ln)
}

func getValidActImages() []ActImageSrc {
	return []ActImageSrc{"https://example.com/image.png"}
}

func getValidActText() ActText {
	return "This act is very useful to enhance your hip joints mobility."
}

func getValidCharacteristics() ActCharacteristics {
	return ActCharacteristics{}
}

func getTooManyActImages() []ActImageSrc {
	images := []ActImageSrc{}
	for i := 0; i < domain.ACT_IMAGES_MAX_NUMBER+1; i++ {
		images = append(images, ActImageSrc("https://example.com/image.png"))
	}
	return images
}

func getTooLongActText() ActText {
	ln := ""
	for i := 0; i < domain.ACT_TEXT_MAX_LENGTH+1; i++ {
		ln += "a"
	}
	return ActText(ln)
}
