package program

import (
	"routine/internal/domain/aggregates/user"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	for _, title := range getInvalidProgramTitleSet() {
		_, err := CreateWeeklyProgram(
			title,
			*user.UserFrom(42),
			nil,
			getValidProgramImageSrcs(),
			getValidProgramText(),
			getValidDailyRoutines(),
		)
		if err != ErrInvalidProgramTitle {
			t.Errorf("creating weekly program with invalid title(%s) should cause ErrInvalidProgramTitle but got %v", title, err)
		}
	}
}

func TestCreateWeeklyProgramSettingVersionsRefs(t *testing.T) {
	np, err := CreateWeeklyProgram(
		getValidProgramTitle(),
		*user.UserFrom(42),
		nil,
		getValidProgramImageSrcs(),
		getValidProgramText(),
		getValidDailyRoutines(),
	)
	if err != nil || np == nil {
		t.Error("expected newProgram to be created but got nil and err:", err)
	}
	for _, v := range np.Versions() {
		if v.Program() != np.Code() {
			t.Errorf("program is expected to set versions' reference to itself but the reference is %v", v.Program())
		}
	}
}

func TestUpgradeWeeklyProgram(t *testing.T) {
	author := user.UserFrom(42)
	np, err := CreateWeeklyProgram(
		getValidProgramTitle(),
		*author,
		nil,
		getValidProgramImageSrcs(),
		getValidProgramText(),
		getValidDailyRoutines(),
	)
	if err != nil || np == nil {
		t.Error("expected newProgram to be created but got nil and err:", err)
	}
	newImageSrcs := ProgramImageSrcs{"https://www.example.com/newimage.png"}
	newText := getValidProgramText() + "new text"
	newDailyRoutines := getValidDailyRoutines()
	newDailyRoutines = newDailyRoutines[:len(newDailyRoutines)-1]
	prevLatestVersion := np.LatestVersion().Version()
	prevVersionsLen := len(np.Versions())
	up, err := np.Upgrade(*author, newImageSrcs, newText, newDailyRoutines)
	if err != nil || up == nil {
		t.Error("expected program to be upgraded but got nil and err:", err)
	}
	if up.LatestVersion().Version() != prevLatestVersion+1 || np.LatestVersion().Version() != prevLatestVersion+1 {
		t.Errorf("expected program to upgrade version but the version is %v", up.LatestVersion().Version())
	}
	if len(up.Versions()) != prevVersionsLen+1 || len(np.Versions()) != prevVersionsLen+1 {
		t.Errorf("expected program to add new version but the versions length is %v", len(up.Versions()))
	}
	if up.LatestVersion().ImageSrcs()[0] != newImageSrcs[0] || np.LatestVersion().ImageSrcs()[0] != newImageSrcs[0] {
		t.Errorf("expected program to upgrade image srcs but the image srcs is %v", up.LatestVersion().ImageSrcs())
	}
	if up.LatestVersion().Text() != newText || np.LatestVersion().Text() != newText {
		t.Errorf("expected program to upgrade text but the text is %v", up.LatestVersion().Text())
	}
	if diff := cmp.Diff(newDailyRoutines, up.LatestVersion().DailyRoutines(), cmp.AllowUnexported(DailyRoutine{}, RoutineAct{}, RoutineActStage{})); diff != "" {
		t.Error("expected program to upgrade daily routines but the daily routines got difference:", diff)
	}
	if diff := cmp.Diff(newDailyRoutines, np.LatestVersion().DailyRoutines(), cmp.AllowUnexported(DailyRoutine{}, RoutineAct{}, RoutineActStage{})); diff != "" {
		t.Error("expected program to upgrade daily routines but the daily routines got difference:", diff)
	}
}
