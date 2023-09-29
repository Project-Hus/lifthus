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
	drs := getValidDailyRoutines()
	v1, err := CreateProgramVersion(
		"ABCDEF12",
		1,
		getValidProgramImageSrcs(),
		getValidProgramText(),
		drs,
	)
	if err != nil || v1 == nil {
		t.Errorf("program version is expected to be created but isn't and got err: %v", err)
	}
	for _, dr := range v1.DailyRoutines() {
		if dr.Version() != v1.Code() {
			t.Errorf("program version is expected to set daily routines' reference to itself but the reference is %v", dr.Version())
		}
	}
}
