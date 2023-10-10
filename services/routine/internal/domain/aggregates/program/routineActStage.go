package program

import "fmt"

func MapRoutineActStage(code string) (RoutineActStage, error) {
	switch code {
	case WARMUP:
		return WarmUpStage, nil
	case MAIN:
		return MainStage, nil
	case COOLDOWN:
		return CoolDownStage, nil
	default:
		return RoutineActStage{}, fmt.Errorf("invalid routine act stage code: %s", code)
	}
}

const (
	WARMUP   = "warmup"
	MAIN     = "main"
	COOLDOWN = "cooldown"
)

var WarmUpStage = RoutineActStage{WARMUP}
var MainStage = RoutineActStage{MAIN}
var CoolDownStage = RoutineActStage{COOLDOWN}

type RoutineActStage struct {
	code string
}

func (ras RoutineActStage) Type() string {
	return ras.code
}
