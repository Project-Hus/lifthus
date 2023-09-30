package program

import (
	"routine/internal/domain"
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

func TestCreateProgramVersionSettingProgramCodeRef(t *testing.T) {
	npv, err := CreateProgramVersion(
		"ABCDEF12",
		2,
		getValidProgramImageSrcs(),
		getValidProgramText(),
		getValidDailyRoutines(),
	)
	if err != nil || npv == nil {
		t.Errorf("program version is expected to be created but isn't and got err: %v", err)
	}
	if npv.Program() != "ABCDEF12" {
		t.Errorf("program version is expected to set program code reference to itself but the reference is %v", npv.Program())
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

func TestCreateProgramVersionFailByImageSrcsConstraints(t *testing.T) {
	_, err := CreateProgramVersion(
		"ABCDEF12",
		42,
		getTooManyProgramImageSrcs(),
		getValidProgramText(),
		getValidDailyRoutines(),
	)
	if err != ErrInvalidProgramImageSrcs {
		t.Errorf("creating program version with invalid image srcs is expected to fail by constraints but got err: %v", err)
	}
}

func TestCreateProgramVersionFailByProgramTextConstraints(t *testing.T) {
	_, err := CreateProgramVersion(
		"ABCDEF12",
		42,
		getValidProgramImageSrcs(),
		getTooShortProgramText(),
		getValidDailyRoutines(),
	)
	if err != ErrInvalidProgramText {
		t.Errorf("creating program version with too short text is expected to fail by constraints but got err: %v", err)
	}
	_, err = CreateProgramVersion(
		"ABCDEF12",
		42,
		getValidProgramImageSrcs(),
		getTooLongProgramText(),
		getValidDailyRoutines(),
	)
	if err != ErrInvalidProgramText {
		t.Errorf("creating program version with too long text is expected to fail by constraints but got err: %v", err)
	}
}
