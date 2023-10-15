package program

import (
	"fmt"
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
	"time"
)

func getValidWeeklyProgramWithAuthor(author user.User) *Program {
	r, err := getValidFirstRelease()
	if err != nil {
		panic(err)
	}
	newProgram, err := CreateProgram(
		WeeklyType,
		getValidProgramTitle(),
		author.Id(),
		nil,
		*r,
	)
	if err != nil {
		panic(err)
	}
	return newProgram
}

func getValidProgramTitle() ProgramTitle {
	title := ProgramTitle("")
	for i := 0; i < domain.PROGRAM_TITLE_MIN_LENGTH; i++ {
		title += "a"
	}
	return title
}

func getValidFirstRelease() (*ProgramRelease, error) {
	return ProgramReleaseFrom(
		ProgramVersionNumber(1),
		domain.CreatedAt(time.Now()),
		getValidProgramImageSrcs(),
		getValidProgramText(),
		getValidRoutines(),
	)
}

func getInvalidProgramTitleSet() []ProgramTitle {
	return []ProgramTitle{
		getTooShortProgramTitle(),
		getTooLongProgramTitle(),
	}
}

func getTooShortProgramTitle() ProgramTitle {
	title := ProgramTitle("")
	for i := 0; i < domain.PROGRAM_TITLE_MIN_LENGTH-1; i++ {
		title += "a"
	}
	return title
}

func getTooLongProgramTitle() ProgramTitle {
	title := ProgramTitle("")
	for i := 0; i < domain.PROGRAM_TITLE_MAX_LENGTH+1; i++ {
		title += "a"
	}
	return title
}

func getValidProgramImageSrcs() ProgramImageSrcs {
	imageSrcs := ProgramImageSrcs{}
	for i := 0; i < domain.PROGRAM_IMAGES_MIN_NUMBER; i++ {
		imageSrcs = append(imageSrcs, "https://www.example.com/image"+fmt.Sprint(i)+".png")
	}
	return imageSrcs
}

func getTooManyProgramImageSrcs() ProgramImageSrcs {
	imageSrcs := ProgramImageSrcs{}
	for i := 0; i < domain.PROGRAM_IMAGES_MAX_NUMBER+1; i++ {
		imageSrcs = append(imageSrcs, "https://www.example.com/image"+fmt.Sprint(i)+".png")
	}
	return imageSrcs
}

func getValidProgramText() ProgramText {
	text := ProgramText("")
	for i := 0; i < domain.PROGRAM_TEXT_MIN_LENGTH; i++ {
		text += "a"
	}
	return text
}

func getTooShortProgramText() ProgramText {
	text := ProgramText("")
	for i := 0; i < domain.PROGRAM_TEXT_MIN_LENGTH-1; i++ {
		text += "a"
	}
	return text
}

func getTooLongProgramText() ProgramText {
	text := ProgramText("")
	for i := 0; i < domain.PROGRAM_TEXT_MAX_LENGTH+1; i++ {
		text += "a"
	}
	return text
}
