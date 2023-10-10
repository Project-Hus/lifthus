package dto

import (
	"routine/internal/domain/aggregates/program"
	"strconv"
	"time"
)

func QueryProgramDtoFrom(p *program.Program) *QueryProgramDto {
	author := strconv.FormatUint(uint64(p.Author()), 10)
	versions := QueryProgramVersionsDtoFrom(p.Versions())
	return &QueryProgramDto{
		Code:        string(p.Code()),
		ProgramType: string(p.ProgramType().Type()),
		Title:       string(p.Title()),
		Author:      author,
		CreatedAt:   time.Time(p.CreatedAt()).String(),
		DerivedFrom: (*string)(p.DerivedFrom()),
		Versions:    versions,
	}
}

type QueryProgramDto struct {
	Code        string                    `json:"code"`
	ProgramType string                    `json:"programType"`
	Title       string                    `json:"title"`
	Author      string                    `json:"author"`
	CreatedAt   string                    `json:"createdAt"`
	DerivedFrom *string                   `json:"derivedFrom"`
	Versions    []*QueryProgramVersionDto `json:"versions"`
}

func QueryProgramVersionsDtoFrom(pvs []*program.ProgramVersion) []*QueryProgramVersionDto {
	qpvs := make([]*QueryProgramVersionDto, len(pvs))
	for i, pv := range pvs {
		qpvs[i] = QueryProgramVersionDtoFrom(pv)
	}
	return qpvs
}

func QueryProgramVersionDtoFrom(pv *program.ProgramVersion) *QueryProgramVersionDto {
	qdrs := QueryDailyRoutinesDtoFrom(pv.DailyRoutines())
	return &QueryProgramVersionDto{
		Code:          string(pv.Code()),
		ProgramCode:   string(pv.Program()),
		Version:       uint(pv.Version()),
		CreatedAt:     time.Time(pv.CreatedAt()).String(),
		ImageSrcs:     pv.ImageSrcs(),
		Text:          string(pv.Text()),
		DailyRoutines: qdrs,
	}
}

type QueryProgramVersionDto struct {
	Code          string                  `json:"code"`
	ProgramCode   string                  `json:"programCode"`
	Version       uint                    `json:"version"`
	CreatedAt     string                  `json:"createdAt"`
	ImageSrcs     []string                `json:"imageSrcs"`
	Text          string                  `json:"text"`
	DailyRoutines []*QueryDailyRoutineDto `json:"dailyRoutines"`
}

func QueryDailyRoutinesDtoFrom(drs []*program.DailyRoutine) []*QueryDailyRoutineDto {
	qdrs := make([]*QueryDailyRoutineDto, len(drs))
	for i, dr := range drs {
		qdrs[i] = QueryDailyRoutineDtoFrom(dr)
	}
	return qdrs
}

func QueryDailyRoutineDtoFrom(dr *program.DailyRoutine) *QueryDailyRoutineDto {
	qras := QueryRoutineActsDtoFrom(dr.RoutineActs())
	return &QueryDailyRoutineDto{
		Code:        string(dr.Code()),
		VersionCode: string(dr.Version()),
		Day:         uint(dr.Day()),
		RoutineActs: qras,
	}
}

type QueryDailyRoutineDto struct {
	Code        string                `json:"code"`
	VersionCode string                `json:"versionCode"`
	Day         uint                  `json:"day"`
	RoutineActs []*QueryRoutineActDto `json:"routineActs"`
}

func QueryRoutineActsDtoFrom(ras []*program.RoutineAct) []*QueryRoutineActDto {
	qras := make([]*QueryRoutineActDto, len(ras))
	for i, ra := range ras {
		qras[i] = QueryRoutineActDtoFrom(ra)
	}
	return qras
}

func QueryRoutineActDtoFrom(ra *program.RoutineAct) *QueryRoutineActDto {
	return &QueryRoutineActDto{
		DailyRoutineCode: string(ra.DailyRoutine()),
		Order:            uint(ra.Order()),
		ActVersionCode:   string(ra.ActVersion()),
		Stage:            ra.Stage().Type(),
		RepsOrMeters:     uint(ra.RepsOrMeters()),
		RatioOrSecs:      float64(ra.RatioOrSecs()),
	}
}

type QueryRoutineActDto struct {
	DailyRoutineCode string  `json:"dailyRoutineCode"`
	Order            uint    `json:"order"`
	ActVersionCode   string  `json:"actVersionCode"`
	Stage            string  `json:"stage"`
	RepsOrMeters     uint    `json:"repsOrMeters"`
	RatioOrSecs      float64 `json:"ratioOrSecs"`
}
