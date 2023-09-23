package program

func TypeFrom(programType programType, iteration uint) ProgramType {
	return ProgramType{
		ptype:     programType.ptype,
		iteration: iteration,
	}
}

const (
	WEEKLY = "weekly"
	DAILY  = "daily"
)

var WeeklyType = programType{WEEKLY}
var DailyType = programType{DAILY}

type programType struct {
	ptype string
}

type ProgramType struct {
	ptype     string
	iteration uint
}
