package program

import (
	"fmt"
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
)

func getValidWeeklyProgramWithAuthor(author user.User) *Program {
	newProgram, err := CreateWeeklyProgram(
		getValidProgramTitle(),
		author,
		nil,
		getValidProgramImageSrcs(),
		getValidProgramText(),
		getValidDailyRoutines(),
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
