package program

import "routine/internal/domain"

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

type DailyRoutines []*DailyRoutine

func (drs DailyRoutines) IsValid() bool {
	if len(drs) == 0 {
		return false
	}
	dCnt := DailyRoutineDay(0)
	for _, dr := range drs {
		if dr.Day() <= dCnt {
			return false
		}
		dCnt = dr.day
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

func (dr DailyRoutine) Version() ProgramVersionCode {
	return dr.version
}

func (dr DailyRoutine) Day() DailyRoutineDay {
	return dr.day
}

func (dr DailyRoutine) RoutineActs() RoutineActs {
	return dr.routineActs
}

func (d *DailyRoutine) setVersion(version ProgramVersionCode) {
	d.version = version
}
