package dto

import (
	"routine/internal/domain/aggregates/act"
	"strconv"
	"time"
)

func QueryActDtoFrom(target *act.Act) *QueryActDto {
	authorStr := strconv.FormatUint(uint64(target.Author()), 10)
	return &QueryActDto{
		Code:      string(target.Code()),
		Author:    authorStr,
		ActType:   target.Type().Type(),
		Name:      string(target.Name()),
		Text:      string(target.Text()),
		ImageSrcs: target.ImageSrcs(),
		CreatedAt: time.Time(target.CreatedAt()).String(),
		Standard:  target.IsStandard(),
	}
}

type QueryActDto struct {
	Code      string   `json:"code"`
	Author    string   `json:"author"`
	ActType   string   `json:"actType"`
	Name      string   `json:"name"`
	Text      string   `json:"text"`
	ImageSrcs []string `json:"imageSrcs"`
	CreatedAt string   `json:"createdAt"`
	Standard  bool     `json:"standard"`
}

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

type UpdateActRequestDto struct {
	ActCode   string    `json:"actCode"`
	Text      *string   `json:"text"`
	ImageSrcs *[]string `json:"imageSrcs"`
}

func (ua UpdateActRequestDto) ToServiceDto() *UpdateActServiceDto {
	return &UpdateActServiceDto{
		ActCode:   ua.ActCode,
		Text:      ua.Text,
		ImageSrcs: ua.ImageSrcs,
	}
}

type UpdateActServiceDto struct {
	ActCode   string
	Text      *string
	ImageSrcs *[]string
}
