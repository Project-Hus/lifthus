package program

import "testing"

func TestCreateRoutineCausingErrWithWrongRoutineActsOrder(t *testing.T) {
	invalidRoutineActsSet := getInvalidRoutineActsSets()
	for _, ras := range invalidRoutineActsSet {
		_, err := CreateRoutine(
			1,
			ras,
		)
		if err != ErrInvalidRoutineActs {
			t.Errorf("invalid routine acts order should cause ErrInvalidRoutineActs but got %v", err)
		}
	}
}

func getValidRoutines() Routines {
	drs := Routines{}
	days := []RoutineDay{3, 4, 11, 19, 21}
	for _, d := range days {
		dr, err := CreateRoutine(
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

func getInvalidRoutinesSets() []Routines {
	drss := []Routines{}
	invalidDaySet := [][]RoutineDay{
		{1, 2, 2, 3, 7},
		{2, 4, 3, 5, 9},
		{0, 1, 2, 3, 4},
		{},
	}
	for _, days := range invalidDaySet {
		drs := Routines{}
		for _, d := range days {
			ras := getValidRoutineActs()
			dr, err := CreateRoutine(d, ras)
			if err != nil {
				panic(err)
			}
			drs = append(drs, dr)
		}
		drss = append(drss, drs)
	}
	return drss
}
