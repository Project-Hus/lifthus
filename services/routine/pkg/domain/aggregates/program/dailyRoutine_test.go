package program

import "testing"

func TestCreateDailyRoutineCausesErrWithWrongRoutineActsOrder(t *testing.T) {
	invalidRoutineActsSet := getInvalidRoutineActsSets()
	for _, ras := range invalidRoutineActsSet {
		_, err := CreateDailyRoutineWithoutProgramVersion(
			1,
			ras,
		)
		if err != ErrInvalidRoutineActs {
			t.Errorf("invalid routine acts order should cause ErrInvalidRoutineActs but got %v", err)
		}
	}
}

func TestCreateDailyRoutineSettingRoutineActsRefs(t *testing.T) {
	ras := getValidRoutineActs()
	dr, err := CreateDailyRoutineWithoutProgramVersion(
		1,
		ras,
	)
	if err != nil || dr == nil {
		t.Errorf("expected daily routine to be created but got nil and err: %v", err)
	}
	for _, ra := range dr.RoutineActs() {
		if ra.DailyRoutine() != dr.Code() {
			t.Errorf("daily routine is expected to set routine acts' reference to itself but got %v", ra.DailyRoutine())
		}
	}
}
