package program

import (
	"routine/pkg/domain/aggregates/user"
	"testing"
)

func TestCreateWeeklyProgram(t *testing.T) {
	newProgram, err := CreateWeeklyProgram(
		getValidProgramTitle(),
		*user.UserFrom(42),
		nil,
		getValidProgramImageSrcs(),
		getValidProgramText(),
		getValidDailyRoutines(),
	)
	if err != nil || newProgram == nil {
		t.Error("expected newProgram to be created but got nil and err:", err)
	}
}

func TestCreateWeeklyProgramFailByTitleConstraints(t *testing.T) {
	t.Errorf("not implemented")
}

func TestUpgradeWeeklyProgram(t *testing.T) {
	t.Errorf("not implemented")
}
