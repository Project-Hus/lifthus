package program

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
