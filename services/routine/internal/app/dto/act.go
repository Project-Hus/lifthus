package dto

import (
	"routine/internal/domain/aggregates/act"
	"strconv"
	"time"
)

type CreateActRequestDto struct {
	ActType   string   `json:"actType"`
	Name      string   `json:"name"`
	Author    string   `json:"author"`
	Text      string   `json:"text"`
	ImageSrcs []string `json:"imageSrcs"`
}

func (ca CreateActRequestDto) ToServiceDto() (*CreateActServiceDto, error) {
	author, err := strconv.ParseUint(ca.Author, 10, 64)
	if err != nil {
		return nil, err
	}
	return &CreateActServiceDto{
		ActType:   ca.ActType,
		Name:      ca.Name,
		Author:    author,
		Text:      ca.Text,
		ImageSrcs: ca.ImageSrcs,
	}, nil
}

type CreateActServiceDto struct {
	ActType   string
	Name      string
	Author    uint64
	Text      string
	ImageSrcs []string
}

type UpgradeActRequestDto struct {
	ActCode   string    `json:"actCode"`
	Text      *string   `json:"text"`
	ImageSrcs *[]string `json:"imageSrcs"`
}

func (ua UpgradeActRequestDto) ToServiceDto() *UpgradeActServiceDto {
	return &UpgradeActServiceDto{
		ActCode:   ua.ActCode,
		Text:      ua.Text,
		ImageSrcs: ua.ImageSrcs,
	}
}

type UpgradeActServiceDto struct {
	ActCode   string
	Text      *string
	ImageSrcs *[]string
}

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
