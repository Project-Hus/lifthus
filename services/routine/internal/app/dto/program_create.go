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

	Routines []CreateProgramRequestRoutineDto `json:"dailyRoutines"`
}

type CreateProgramRequestRoutineDto struct {
	Day         int                                 `json:"day"`
	RoutineActs []CreateProgramRequestRoutineActDto `json:"routineActs"`
}

type CreateProgramRequestRoutineActDto struct {
	Order        int     `json:"order"`
	ActCode      string  `json:"actCode"`
	Stage        string  `json:"stage"`
	RepsOrMeters int     `json:"repsOrMeters"`
	RatioOrSecs  float64 `json:"ratioOrSecs"`
}

func (cpr CreateProgramRequestDto) ToServiceDto() (*CreateProgramServiceDto, error) {
	author, err := strconv.ParseInt(cpr.Author, 10, 64)
	if err != nil {
		return nil, err
	}

	svcDRs := make([]CreateProgramServiceRoutineDto, len(cpr.Routines))
	for i, rdr := range cpr.Routines {
		svcRAs := make([]CreateProgramServiceRoutineActDto, len(rdr.RoutineActs))
		for j, rra := range rdr.RoutineActs {
			svcRAs[j] = CreateProgramServiceRoutineActDto(rra)
		}
		svcDRs[i] = CreateProgramServiceRoutineDto{
			Day:         rdr.Day,
			RoutineActs: svcRAs,
		}
	}

	return &CreateProgramServiceDto{
		ProgramType: cpr.ProgramType,
		Title:       cpr.Title,
		Author:      author,
		DerivedFrom: cpr.DerivedFrom,
		ImageSrcs:   cpr.ImageSrcs,
		Text:        cpr.Text,
		Routines:    svcDRs,
	}, nil
}

type CreateProgramServiceDto struct {
	ProgramType string
	Title       string

	Author      int64
	DerivedFrom *string

	ImageSrcs []string
	Text      string

	Routines []CreateProgramServiceRoutineDto
}

type CreateProgramServiceRoutineDto struct {
	Day         int
	RoutineActs []CreateProgramServiceRoutineActDto
}

type CreateProgramServiceRoutineActDto struct {
	Order        int
	ActCode      string
	Stage        string
	RepsOrMeters int
	RatioOrSecs  float64
}
