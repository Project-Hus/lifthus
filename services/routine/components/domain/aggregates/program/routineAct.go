package program

import "routine/components/domain"

type RoutineActId domain.Id
type RoutineActOrder uint

type RoutineAct struct {
	id    *RoutineActId
	order RoutineActOrder
	//actId ActId
}
