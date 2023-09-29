package program

import "routine/pkg/domain"

func CreateDailyRoutine(
	version ProgramVersionCode,
	day DailyRoutineDay,
	routineActs []*RoutineAct,
) (*DailyRoutine, error) {
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	return DailyRoutineFrom(DailyRoutineCode(code), version, day, routineActs)
}

func DailyRoutineFrom(
	code DailyRoutineCode,
	version ProgramVersionCode,
	day DailyRoutineDay,
	routineActs RoutineActs,
) (*DailyRoutine, error) {
	if !routineActs.IsValid() {
		return nil, ErrInvalidRoutineActs
	}
	return &DailyRoutine{
		code:        code,
		version:     version,
		day:         day,
		routineActs: routineActs,
	}, nil
}

type DailyRoutineCode domain.Code
type DailyRoutineDay uint
type RoutineActs []*RoutineAct

func (ras RoutineActs) IsValid() bool {
	for i, ra := range ras {
		if i != int(ra.order)-1 {
			return false
		}
	}
	return true
}

type DailyRoutine struct {
	code DailyRoutineCode

	version ProgramVersionCode
	day     DailyRoutineDay

	routineActs RoutineActs
}
