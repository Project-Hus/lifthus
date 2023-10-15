package program

func CreateRoutine(
	day RoutineDay,
	routineActs RoutineActs,
) (*Routine, error) {
	return RoutineFrom(day, routineActs)
}

func RoutineFrom(
	day RoutineDay,
	routineActs RoutineActs,
) (*Routine, error) {
	if !routineActs.IsValid() {
		return nil, ErrInvalidRoutineActs
	}
	return &Routine{
		day:         day,
		routineActs: routineActs,
	}, nil
}

type RoutineDay int

type Routines []*Routine

func (drs Routines) IsValid() bool {
	if len(drs) == 0 {
		return false
	}
	dCnt := RoutineDay(0)
	for _, dr := range drs {
		if dr.Day() <= dCnt {
			return false
		}
		dCnt = dr.day
	}
	return true
}

type Routine struct {
	day         RoutineDay
	routineActs RoutineActs
}

func (dr Routine) Day() RoutineDay {
	return dr.day
}

func (dr Routine) RoutineActs() RoutineActs {
	return dr.routineActs
}
