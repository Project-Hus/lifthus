package program

type ProgramType struct {
	programType string
}

const (
	WEEKLY = "weekly"
	DAILY  = "daily"
)

var WeeklyType = ProgramType{WEEKLY}
var DailyType = ProgramType{DAILY}

func (pt ProgramType) Type() string {
	return pt.programType
}
