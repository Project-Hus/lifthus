package program

import (
	"log"
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"time"
)

func isCreationSuccessWith(
	md programMetadata,
	drv programDerivations,
	inf programInfo,
) bool {
	_, err := CreateProgram(md, drv, inf)
	return err == nil
}

func getNormalTitle() ProgramTitle {
	return "Normal Title"
}

func getOverflownTitle() ProgramTitle {
	title := ""
	for i := 0; i < TITLE_MAX_LENGTH+1; i++ {
		title += "*"
	}
	return ProgramTitle(title)
}

func getNormalImageSrcs() ProgramImageSrcs {
	return []string{"https://www.google.com"}
}

func getOverflownImageSrcs() ProgramImageSrcs {
	imageSrcs := []string{}
	for i := 0; i < IMAGES_MAX_COUNT+1; i++ {
		imageSrcs = append(imageSrcs, "https://www.google.com")
	}
	return imageSrcs
}

func getNormalDescription() ProgramDescription {
	return "What doens't kill you makes you stronger."
}

func getOverflownDescription() ProgramDescription {
	description := ""
	for i := 0; i < DESCRIPTION_MAX_LENGTH+1; i++ {
		description += "*"
	}
	return ProgramDescription(description)
}

func getArgsToSucceedCreateProgram() (programMetadata, programDerivations, programInfo) {
	return NewProgramMetadata(42, WeeklyType, domain.CreatedAt(time.Now())),
		ProgramDerivationsFrom(nil, []ProgramId{}),
		ProgramInfoFrom("Lifthus Program", []string{}, "this program is so cool")
}

var testProgramId ProgramId = 42

func getProgramWithAuthor(authorId user.UserId) *Program {
	code, err := domain.RandomHexCode()
	if err != nil {
		log.Println("RandomHexCode() failed")
	}
	newProgram := ProgramFrom(
		testProgramId,
		ProgramCode(code),
		ProgramMetadataFrom(authorId, WeeklyType, domain.CreatedAt(time.Now()), nil),
		ProgramDerivationsFrom(nil, []ProgramId{}),
		ProgramInfoFrom(getNormalTitle(), getNormalImageSrcs(), getNormalDescription()),
	)
	testProgramId++
	return newProgram
}

func getProgramWithDerivingProgramsAndAuthor() (*Program, *user.User) {
	author := user.UserFrom(42)
	newProgram := getProgramWithAuthor(author.Id())
	newProgram.deriving = []ProgramId{43, 44}
	return newProgram, author
}
