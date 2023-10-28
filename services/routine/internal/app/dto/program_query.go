package dto

import (
	"routine/internal/domain/aggregates/program"
	"strconv"
	"time"
)

func QueryProgramDtoFrom(p *program.Program) *QueryProgramDto {
	author := strconv.FormatInt(int64(p.Author()), 10)
	versions := QueryProgramReleasesDtoFrom(p.Releases())
	ppv := p.ParentProgramVersion()
	var ppc *string
	var pv *int
	if ppv != nil {
		ppc = (*string)(&ppv.ProgramCode)
		pv = (*int)(&ppv.ProgramVersionNumber)
	}
	return &QueryProgramDto{
		Code:              string(p.Code()),
		ProgramType:       string(p.ProgramType().Type()),
		Title:             string(p.Title()),
		Author:            author,
		CreatedAt:         time.Time(p.CreatedAt()).String(),
		ParentProgramCode: ppc,
		ParentVersion:     pv,
		Releases:          versions,
	}
}

type QueryProgramDto struct {
	Code              string                    `json:"code"`
	ProgramType       string                    `json:"programType"`
	Title             string                    `json:"title"`
	Author            string                    `json:"author"`
	CreatedAt         string                    `json:"createdAt"`
	ParentProgramCode *string                   `json:"parentProgramCode,omitempty"`
	ParentVersion     *int                      `json:"parentVersion,omitempty"`
	Releases          []*QueryProgramReleaseDto `json:"releases"`
}

func QueryProgramReleasesDtoFrom(pvs []*program.ProgramRelease) []*QueryProgramReleaseDto {
	qpvs := make([]*QueryProgramReleaseDto, len(pvs))
	for i, pv := range pvs {
		qpvs[i] = QueryProgramReleaseDtoFrom(pv)
	}
	return qpvs
}

func QueryProgramReleaseDtoFrom(pv *program.ProgramRelease) *QueryProgramReleaseDto {
	qdrs := QueryRoutinesDtoFrom(pv.Routines())
	return &QueryProgramReleaseDto{
		Version:   int(pv.Version()),
		CreatedAt: time.Time(pv.CreatedAt()).String(),
		Text:      string(pv.Text()),
		ImageSrcs: pv.ImageSrcs(),
		Routines:  qdrs,
	}
}

type QueryProgramReleaseDto struct {
	Version   int                `json:"version"`
	CreatedAt string             `json:"createdAt"`
	Text      string             `json:"text"`
	ImageSrcs []string           `json:"imageSrcs"`
	Routines  []*QueryRoutineDto `json:"routines"`
}

func QueryRoutinesDtoFrom(drs []*program.Routine) []*QueryRoutineDto {
	qdrs := make([]*QueryRoutineDto, len(drs))
	for i, dr := range drs {
		qdrs[i] = QueryRoutineDtoFrom(dr)
	}
	return qdrs
}

func QueryRoutineDtoFrom(dr *program.Routine) *QueryRoutineDto {
	qras := QueryRoutineActsDtoFrom(dr.RoutineActs())
	return &QueryRoutineDto{
		Day:         int(dr.Day()),
		RoutineActs: qras,
	}
}

type QueryRoutineDto struct {
	Day         int                   `json:"day"`
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
		Order:        int(ra.Order()),
		ActCode:      string(ra.Act()),
		Stage:        ra.Stage().Type(),
		RepsOrMeters: int(ra.RepsOrMeters()),
		RatioOrSecs:  float64(ra.RatioOrSecs()),
	}
}

type QueryRoutineActDto struct {
	Order        int     `json:"order"`
	ActCode      string  `json:"actCode"`
	Stage        string  `json:"stage"`
	RepsOrMeters int     `json:"repsOrMeters"`
	RatioOrSecs  float64 `json:"ratioOrSecs"`
}
