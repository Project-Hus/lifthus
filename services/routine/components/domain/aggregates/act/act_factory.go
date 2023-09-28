package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"time"
)

func CreateWeightAct(name ActName, author user.User, desc ActDescription) (*Act, error) {
	if !IsActNameValid(name) || !IsActImagesValid(desc.ImageSrcs) || !IsActTextValid(desc.Text) {
		return nil, ErrInvalidActInfo
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	return &Act{
		id:              nil,
		code:            ActCode(code),
		actType:         WeightType,
		name:            name,
		version:         1,
		author:          author.Id(),
		createdAt:       domain.CreatedAt(time.Now()),
		updatedAt:       nil,
		imageSrcs:       desc.ImageSrcs,
		text:            desc.Text,
		characteristics: desc.Characteristics,
	}, nil
}

func CreatTimeAct(name ActName, author user.User, desc ActDescription) (*Act, error) {
	if !IsActNameValid(name) || !IsActImagesValid(desc.ImageSrcs) || !IsActTextValid(desc.Text) {
		return nil, ErrInvalidActInfo
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	return &Act{
		id:              nil,
		code:            ActCode(code),
		actType:         TimeType,
		name:            name,
		version:         1,
		author:          author.Id(),
		createdAt:       domain.CreatedAt(time.Now()),
		updatedAt:       nil,
		imageSrcs:       desc.ImageSrcs,
		text:            desc.Text,
		characteristics: desc.Characteristics,
	}, nil
}

func CreateSimpleAct(name ActName, author user.User, desc ActDescription) (*Act, error) {
	if !IsActNameValid(name) || !IsActImagesValid(desc.ImageSrcs) || !IsActTextValid(desc.Text) {
		return nil, ErrInvalidActInfo
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	return &Act{
		id:              nil,
		code:            ActCode(code),
		actType:         SimpleType,
		name:            name,
		version:         1,
		author:          author.Id(),
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
