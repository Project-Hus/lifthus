package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"time"
)

func CreateAct(base ActBase, author user.UserId, desc ActDescription) (*Act, error) {
	if !IsNewActValid(base, desc) {
		return nil, ErrInvalidActInfo
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	return &Act{
		id:              nil,
		code:            ActCode(code),
		actType:         base.ActType,
		name:            base.ActName,
		version:         1,
		author:          author,
		createdAt:       domain.CreatedAt(time.Now()),
		updatedAt:       nil,
		imageSrcs:       desc.ImageSrcs,
		text:            desc.Text,
		characteristics: desc.Characteristics,
	}, nil
}

func ActFrom(id ActId, code ActCode, base ActBase, md ActMetadata, desc ActDescription) *Act {
	return &Act{
		id:              &id,
		code:            code,
		actType:         base.ActType,
		name:            base.ActName,
		version:         base.ActVersion,
		author:          base.Author,
		createdAt:       md.CreatedAt,
		updatedAt:       md.UpdatedAt,
		imageSrcs:       desc.ImageSrcs,
		text:            desc.Text,
		characteristics: desc.Characteristics,
	}
}
