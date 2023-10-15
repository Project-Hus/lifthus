package program

import "fmt"

const (
	WEEKLY = "weekly"
	DAILY  = "daily"
)

func MapProgramType(pt string) (*ProgramType, error) {
	switch pt {
	case WEEKLY:
		return &WeeklyType, nil
	case DAILY:
		return &DailyType, nil
	default:
		return nil, fmt.Errorf("unknown program type: %s", pt)
	}
}

var WeeklyType = ProgramType{WEEKLY}
var DailyType = ProgramType{DAILY}

type ProgramType struct {
	code string
}

func (pt ProgramType) Type() string {
	return pt.code
}
