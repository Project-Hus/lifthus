package program

import (
	"routine/pkg/domain"
	"testing"
	"time"
)

func TestProgramVersionCausingErrWithInvalidDailyRoutinesOrder(t *testing.T) {
	invalidDailyRoutinesSets := getInvalidDailyRoutinesSets()
	for _, dailyRoutines := range invalidDailyRoutinesSets {
		_, err := ProgramVersionFrom(
			"ABCDEF12",
			"ABCDEF12",
			1,
			domain.CreatedAt(time.Now()),
			getValidProgramImageSrcs(),
			getValidProgramText(),
			dailyRoutines,
		)
		if err != ErrInvalidDailyRoutines {
			t.Errorf("ProgramVersion instantiation with daily routines in invalid order should cause ErrInvalidDailyRoutines but got %v", err)
		}
	}
}

func TestCreateProgramVersionSettingDailyRoutinesRefs(t *testing.T) {
	t.Error("not implemented")
}
