package program

import (
	"routine/internal/domain/aggregates/user"
	"testing"
)

func TestCreateWeeklyProgram(t *testing.T) {
	r, err := getValidFirstRelease()
	if err != nil {
		panic(err)
	}
	newProgram, err := CreateProgram(
		WeeklyType,
		getValidProgramTitle(),
		user.UserFrom(42).Id(),
		nil,
		*r,
	)
	if err != nil || newProgram == nil {
		t.Error("expected newProgram to be created but got nil and err:", err)
	}
}

func TestCreateWeeklyProgramFailByTitleConstraints(t *testing.T) {
	for _, title := range getInvalidProgramTitleSet() {
		r, err := getValidFirstRelease()
		if err != nil {
			panic(err)
		}
		_, err = CreateProgram(
			WeeklyType,
			title,
			user.UserFrom(42).Id(),
			nil,
			*r,
		)
		if err != ErrInvalidProgramTitle {
			t.Errorf("creating weekly program with invalid title(%s) should cause ErrInvalidProgramTitle but got %v", title, err)
		}
	}
}
