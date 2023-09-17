package domain

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

type ProgramContents struct {
	programType ProgramType
	iteration   int
}

func ProgramContentsFrom(
	pt ProgramType,
	itr int,
) ProgramContents {
	return ProgramContents{
		programType: pt,
		iteration:   itr,
	}
}

func (pc ProgramContents) Type() ProgramType {
	return pc.programType
}

func (pc ProgramContents) Iteration() int {
	return pc.iteration
}
