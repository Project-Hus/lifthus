package actversion

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/act"
	"time"
)

func CreateActVersion(act act.Act) *ActVersion {
	ab := act.Base()
	amd := act.Metadata()
	ad := act.Description()
	createdAt := amd.CreatedAt
	if amd.UpdatedAt != nil {
		createdAt = domain.CreatedAt(time.Time(*amd.UpdatedAt))
	}
	return &ActVersion{
		code:    act.Code(),
		version: ab.Version,

		actType: ab.ActType,
		name:    ab.Name,
		author:  ab.Author,

		createdAt: createdAt,

		imageSrcs:       ad.ImageSrcs,
		text:            ad.Text,
		characteristics: ad.Characteristics,
	}
}

func ActVersionFrom(code act.ActCode, version act.ActVersion, avb ActVersionBase, createdAt domain.CreatedAt, desc act.ActDescription) *ActVersion {
	return &ActVersion{
		code:    code,
		version: version,

		actType: avb.ActType,
		name:    avb.Name,
		author:  avb.Author,

		createdAt: createdAt,

		imageSrcs:       desc.ImageSrcs,
		text:            desc.Text,
		characteristics: desc.Characteristics,
	}
}
