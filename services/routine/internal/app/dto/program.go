package dto

import (
	"strconv"
)

type CreateProgramRequestDto struct {
	ProgramType string  `json:"programType"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	DerivedFrom *string `json:"derivedFrom"`

	ImageSrcs []string `json:"imageSrcs"`
	Text      string   `json:"text"`

	DailyRoutines []CreateProgramRequestDailyRoutineDto
}

type CreateProgramRequestDailyRoutineDto struct {
	Day         uint                                `json:"day"`
	RoutineActs []CreateProgramRequestRoutineActDto `json:"routineActs"`
}

type CreateProgramRequestRoutineActDto struct {
	Order        uint    `json:"order"`
	ActVersion   string  `json:"actVersion"`
	Stage        string  `json:"stage"`
	RepsOrMeters uint    `json:"repsOrMeters"`
	RatioOrSecs  float64 `json:"ratioOrSecs"`
}

func (cpr CreateProgramRequestDto) ToServiceDto() (*CreateProgramServiceDto, error) {
	author, err := strconv.ParseUint(cpr.Author, 10, 64)
	if err != nil {
		return nil, err
	}

	svcDRs := make([]CreateProgramServiceDailyRoutineDto, len(cpr.DailyRoutines))
	for i, rdr := range cpr.DailyRoutines {
		svcRAs := make([]CreateProgramServiceRoutineActDto, len(rdr.RoutineActs))
		for j, rra := range rdr.RoutineActs {
			svcRAs[j] = CreateProgramServiceRoutineActDto(rra)
		}
		svcDRs[i] = CreateProgramServiceDailyRoutineDto{
			Day:         rdr.Day,
			RoutineActs: svcRAs,
		}
	}

	return &CreateProgramServiceDto{
		Author:        author,
		DerivedFrom:   cpr.DerivedFrom,
		ImageSrcs:     cpr.ImageSrcs,
		Text:          cpr.Text,
		DailyRoutines: svcDRs,
	}, nil
}

type CreateProgramServiceDto struct {
	Author      uint64
	DerivedFrom *string

	ImageSrcs []string
	Text      string

	DailyRoutines []CreateProgramServiceDailyRoutineDto
}

type CreateProgramServiceDailyRoutineDto struct {
	Day         uint
	RoutineActs []CreateProgramServiceRoutineActDto
}

type CreateProgramServiceRoutineActDto struct {
	Order        uint
	ActVersion   string
	Stage        string
	RepsOrMeters uint
	RatioOrSecs  float64
}
