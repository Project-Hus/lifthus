package program

import (
	"log"
	"routine/pkg/domain"
	"routine/pkg/domain/aggregates/user"
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
		log.Fatal(err)
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
		imageSrcs = append(imageSrcs, "https://www.example.com/image.png")
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

func getValidDailyRoutines() DailyRoutines {
	drs := DailyRoutines{}
	days := []DailyRoutineDay{3, 4, 11, 19, 21}
	for _, d := range days {
		dr, err := CreateDailyRoutineWithoutProgramVersion(
			d,
			getValidRoutineActs(),
		)
		if err != nil {
			log.Fatal(err)
		}
		drs = append(drs, dr)
	}
	return drs
}
