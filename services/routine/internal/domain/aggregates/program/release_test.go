package program

import (
	"routine/internal/domain"
	"testing"
	"time"
)

func TestProgramReleaseCausingErrWithInvalidRoutinesOrder(t *testing.T) {
	invalidRoutinesSets := getInvalidRoutinesSets()
	for _, dailyRoutines := range invalidRoutinesSets {
		_, err := ProgramReleaseFrom(
			1,
			domain.CreatedAt(time.Now()),
			getValidProgramImageSrcs(),
			getValidProgramText(),
			dailyRoutines,
		)
		if err != ErrInvalidRoutines {
			t.Errorf("ProgramRelease instantiation with daily routines in invalid order should cause ErrInvalidRoutines but got %v", err)
		}
	}
}

func TestCreateProgramReleaseFailByImageSrcsConstraints(t *testing.T) {
	_, err := CreateProgramRelease(
		42,
		getTooManyProgramImageSrcs(),
		getValidProgramText(),
		getValidRoutines(),
	)
	if err != ErrInvalidProgramImageSrcs {
		t.Errorf("creating program version with invalid image srcs is expected to fail by constraints but got err: %v", err)
	}
}

func TestCreateProgramReleaseFailByProgramTextConstraints(t *testing.T) {
	_, err := CreateProgramRelease(
		42,
		getValidProgramImageSrcs(),
		getTooShortProgramText(),
		getValidRoutines(),
	)
	if err != ErrInvalidProgramText {
		t.Errorf("creating program version with too short text is expected to fail by constraints but got err: %v", err)
	}
	_, err = CreateProgramRelease(
		42,
		getValidProgramImageSrcs(),
		getTooLongProgramText(),
		getValidRoutines(),
	)
	if err != ErrInvalidProgramText {
		t.Errorf("creating program version with too long text is expected to fail by constraints but got err: %v", err)
	}
}
