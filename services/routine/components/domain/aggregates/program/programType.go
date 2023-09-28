package program

const (
	WEEKLY = "weekly"
	DAILY  = "daily"
)

var WeeklyType = ProgramType{WEEKLY}
var DailyType = ProgramType{DAILY}

type ProgramType struct {
	code string
}

func (pt ProgramType) Type() string {
	return pt.code
}
