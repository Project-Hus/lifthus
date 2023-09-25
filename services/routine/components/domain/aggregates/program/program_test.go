package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"testing"
)

func TestCreateProgramFailByTitle(t *testing.T) {
	md, drv, _ := getArgsToSucceedCreateProgram()
	info := ProgramInfoFrom(getOverflownTitle(), getNormalImageSrcs(), getNormalDescription())
	if isCreationSuccessWith(md, drv, info) {
		t.Errorf("title longer than max length should fail, but succeeded")
	}
}

func TestCreateProgramFailByImageSrcs(t *testing.T) {
	md, drv, _ := getArgsToSucceedCreateProgram()
	info := ProgramInfoFrom(getNormalTitle(), getOverflownImageSrcs(), getNormalDescription())
	if isCreationSuccessWith(md, drv, info) {
		t.Errorf("images more than max count should fail, but succeeded")
	}
}

func TestCreateProgramFailByDescription(t *testing.T) {
	md, drv, _ := getArgsToSucceedCreateProgram()
	info := ProgramInfoFrom(getNormalTitle(), getNormalImageSrcs(), getOverflownDescription())
	if isCreationSuccessWith(md, drv, info) {
		t.Errorf("description longer than max length should fail, but succeeded")
	}
}

func TestUpdateProgram(t *testing.T) {
	author := user.UserFrom(42)
	program := getProgramWithAuthor(author.Id())
	newTitle := ProgramTitle("Newly Updated Program Title")
	updatedProgramReference, _ := program.Update(author, ProgramUpdates{Title: &newTitle})
	if program.title != newTitle {
		t.Errorf("title should be updated, but not")
	}
	if updatedProgramReference.title != newTitle {
		t.Errorf("title should be updated, but not")
	}
}

func TestUpdateProgramFailByUnauthorizedUpdater(t *testing.T) {
	author := user.UserFrom(42)
	program := getProgramWithAuthor(author.Id())
	abuser := user.UserFrom(43)
	newTitle := ProgramTitle("Newly Updated Program Title")
	_, err := program.Update(abuser, ProgramUpdates{Title: &newTitle})
	if err != domain.ErrUnauthorized {
		t.Errorf("should fail by unauthorized updater, but not")
	}
}

func TestDeleteProgram(t *testing.T) {
	author := user.UserFrom(42)
	program := getProgramWithAuthor(author.Id())
	deletedProgram, _ := program.Delete(author)
	if deletedProgram == nil {
		t.Errorf("should not be nil, but nil")
	}
}

func TestDeleteProgramFailByUnauthorizedDeleter(t *testing.T) {
	author := user.UserFrom(42)
	program := getProgramWithAuthor(author.Id())
	abuser := user.UserFrom(43)
	_, err := program.Delete(abuser)
	if err != domain.ErrUnauthorized {
		t.Errorf("should fail by unauthorized deleter, but not")
	}
}
