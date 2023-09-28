package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"time"
)

func getValidActDescription() ActDescription {
	return ActDescriptionFrom(getValidActImages(), getValidActText(), getValidCharacteristics())
}

func getValidActName() ActName {
	an := ""
	for i := 0; i < domain.ACT_NAME_MIN_LENGTH; i++ {
		an += "a"
	}
	return ActName(an)
}

func getTooLongActName() ActName {
	ln := ""
	for i := 0; i < domain.ACT_NAME_MAX_LENGTH+1; i++ {
		ln += "a"
	}
	return ActName(ln)
}

func getValidActImages() []ActImageSrc {
	ai := []ActImageSrc{}
	for i := 0; i < domain.ACT_IMAGES_MIN_NUMBER+1; i++ {
		ai = append(ai, ActImageSrc("https://example.com/image.png"))
	}
	return ai
}

func getValidActText() ActText {
	at := ""
	for i := 0; i < domain.ACT_TEXT_MIN_LENGTH; i++ {
		at += "a"
	}
	return ActText(at)
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

func getValidActWithAuthor(author user.User) *Act {
	code, _ := domain.RandomHexCode()
	return ActFrom(ActCode(code), getValidActBaseWithAuthor(author), getValidActMetadata(), getValidActDescription())
}

func getValidActBaseWithAuthor(author user.User) ActBase {
	return ActBaseFrom(WeightType, getValidActName(), 42, author.Id())
}

func getValidActMetadata() ActMetadata {
	return ActMetadataFrom(domain.CreatedAt(time.Now()), nil)
}
