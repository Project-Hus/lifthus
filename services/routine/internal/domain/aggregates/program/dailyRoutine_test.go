package program

import "testing"

func TestCreateDailyRoutineCausingErrWithWrongRoutineActsOrder(t *testing.T) {
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
			t.Errorf("daily routine is expected to set routine acts' reference to itself but the reference is %v", ra.DailyRoutine())
		}
	}
}

func getValidDailyRoutines() DailyRoutines {
	drs := DailyRoutines{}
	days := []DailyRoutineDay{3, 4, 11, 19, 21}
	for _, d := range days {
		dr, err := CreateDailyRoutineWithoutProgramVersion(
			d,
			getValidRoutineActs(),
		)
		if err != nil {
			panic(err)
		}
		drs = append(drs, dr)
	}
	return drs
}

func getInvalidDailyRoutinesSets() []DailyRoutines {
	drss := []DailyRoutines{}
	invalidDaySet := [][]DailyRoutineDay{
		{1, 2, 2, 3, 7},
		{2, 4, 3, 5, 9},
		{0, 1, 2, 3, 4},
		{},
	}
	for _, days := range invalidDaySet {
		drs := DailyRoutines{}
		for _, d := range days {
			ras := getValidRoutineActs()
			dr, err := CreateDailyRoutineWithoutProgramVersion(d, ras)
			if err != nil {
				panic(err)
			}
			drs = append(drs, dr)
		}
		drss = append(drss, drs)
	}
	return drss
}
