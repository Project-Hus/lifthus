package program

const (
	WEEKLY = "weekly"
	DAILY  = "daily"
)

var WeeklyType = ProgramType{WEEKLY}
var DailyType = ProgramType{DAILY}

type ProgramType struct {
	programType string
}

func (pt ProgramType) Type() string {
	return pt.programType
}
