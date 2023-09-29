package program

import "routine/pkg/domain"

func CreateDailyRoutineWithoutProgramVersion(
	day DailyRoutineDay,
	routineActs RoutineActs,
) (*DailyRoutine, error) {
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	setDailyRoutineRefToRoutineActs(DailyRoutineCode(code), routineActs)
	return DailyRoutineFrom(DailyRoutineCode(code), ProgramVersionCode(""), day, routineActs)
}

func setDailyRoutineRefToRoutineActs(code DailyRoutineCode, ras RoutineActs) {
	for _, ra := range ras {
		ra.setDailyRoutine(code)
	}
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
	if len(ras) == 0 {
		return false
	}
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

func (dr DailyRoutine) Code() DailyRoutineCode {
	return dr.code
}

func (dr DailyRoutine) Day() DailyRoutineDay {
	return dr.day
}

func (dr DailyRoutine) RoutineActs() RoutineActs {
	return dr.routineActs
}
