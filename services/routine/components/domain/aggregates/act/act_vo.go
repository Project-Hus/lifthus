package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

func ActBaseFrom(
	actType ActType,
	actName ActName,
	actVersion ActVersion,
	author user.UserId,
) ActBase {
	return ActBase{
		ActType:    actType,
		ActName:    actName,
		ActVersion: actVersion,
		Author:     author,
	}
}

type ActBase struct {
	ActType    ActType
	ActName    ActName
	ActVersion ActVersion
	Author     user.UserId
}

func ActMetadataFrom(
	createdAt domain.CreatedAt,
	updatedAt *domain.UpdatedAt,
) ActMetadata {
	return ActMetadata{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type ActMetadata struct {
	CreatedAt domain.CreatedAt
	UpdatedAt *domain.UpdatedAt
}

func ActDescriptionFrom(
	imageSrcs []ActImageSrc,
	text ActText,
	characteristics ActCharacteristics,
) ActDescription {
	return ActDescription{
		ImageSrcs:       imageSrcs,
		Text:            text,
		Characteristics: characteristics,
	}
}

type ActDescription struct {
	ImageSrcs       []ActImageSrc
	Text            ActText
	Characteristics ActCharacteristics
}

type ActUpdates struct {
	ImageSrcs       *[]ActImageSrc
	Text            *ActText
	Characteristics *ActCharacteristics
}

type ActUpdateTargets struct {
	ImageSrcs       []ActImageSrc
	Text            ActText
	Characteristics ActCharacteristics
}
