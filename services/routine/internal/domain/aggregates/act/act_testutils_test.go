package act

import (
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
	"time"
)

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

func getValidActImages() ActImageSrcs {
	ai := ActImageSrcs{}
	for i := 0; i < domain.ACT_IMAGES_MIN_NUMBER+1; i++ {
		ai = append(ai, "https://example.com/image.png")
	}
	return ai
}

func getTooManyActImages() ActImageSrcs {
	ai := ActImageSrcs{}
	for i := 0; i < domain.ACT_IMAGES_MAX_NUMBER+1; i++ {
		ai = append(ai, "https://example.com/image.png")
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

func getTooLongActText() ActText {
	ln := ""
	for i := 0; i < domain.ACT_TEXT_MAX_LENGTH+1; i++ {
		ln += "a"
	}
	return ActText(ln)
}

func getValidActWithAuthor(author user.User) *Act {
	return &Act{
		code:      "ABCDEF12",
		author:    author.Id(),
		actType:   SimpleType,
		name:      getValidActName(),
		text:      getValidActText(),
		imageSrcs: getValidActImages(),
		createdAt: domain.CreatedAt(time.Now()),
		standard:  false,
	}
}
