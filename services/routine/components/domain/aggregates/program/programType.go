package program

const (
	WEEKLY = "weekly"
	DAILY  = "daily"
)

var WeeklyType = ProgramType{WEEKLY}
var DailyType = ProgramType{DAILY}

type ProgramType struct {
	ptype string
}

func (pt ProgramType) Type() string {
	return pt.ptype
}
