package dto

import (
	"os"
	"routine/internal/domain/aggregates/act"
	"strconv"
	"time"
)

type CreateActRequestDto struct {
	ActType string    `json:"actType"`
	Name    string    `json:"name"`
	Author  uint64    `json:"author"`
	Text    string    `json:"text"`
	Images  []os.File `json:"images"`
}

type CreateActDto struct {
	ActType   string `json:"actType"`
	Name      string `json:"name"`
	Author    uint64 `json:"author"`
	Text      string `json:"text"`
	ImageSrcs []string
}

type UpgradeActDto struct{}

func QueryActDtoFrom(target *act.Act) *QueryActDto {
	authorStr := strconv.FormatUint(uint64(target.Author()), 10)
	vdtos := QueryVersionsDtoFrom(target.Versions())
	return &QueryActDto{
		Code:      string(target.Code()),
		ActType:   target.Type().Type(),
		Name:      string(target.Name()),
		Author:    authorStr,
		CreatedAt: time.Time(target.CreatedAt()).Format(time.RFC3339),
		Versions:  vdtos,
	}
}

type QueryActDto struct {
	Code string `json:"code"`

	ActType   string `json:"actType"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	CreatedAt string `json:"createdAt"`

	Versions []QueryVersionDto `json:"versions"`
}

func QueryVersionsDtoFrom(versions []*act.ActVersion) []QueryVersionDto {
	qvdtos := make([]QueryVersionDto, len(versions))
	for i, v := range versions {
		qvdtos[i] = QueryVersionDtoFrom(v)
	}
	return qvdtos
}

func QueryVersionDtoFrom(version *act.ActVersion) QueryVersionDto {
	return QueryVersionDto{
		Code:      string(version.Code()),
		Version:   float64(version.Version()),
		ImageSrcs: version.ImageSrcs(),
		Text:      string(version.Text()),
		CreatedAt: time.Time(version.CreatedAt()).Format(time.RFC3339),
	}
}

type QueryVersionDto struct {
	Code      string   `json:"code"`
	Version   float64  `json:"version"`
	ImageSrcs []string `json:"imageSrcs"`
	Text      string   `json:"text"`
	CreatedAt string   `json:"createdAt"`
}
