package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

func NewProgramMetadata(
	title ProgramTitle,
	author user.UserId,
	programType ProgramType,
	createdAt domain.CreatedAt,
) programMetadata {
	return ProgramMetadataFrom(title, author, programType, createdAt, nil)
}

func ProgramMetadataFrom(
	title ProgramTitle,
	author user.UserId,
	programType ProgramType,
	createdAt domain.CreatedAt,
	updatedAt *domain.UpdatedAt,
) programMetadata {
	return programMetadata{
		Author:      author,
		ProgramType: programType,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

type programMetadata struct {
	Title       ProgramTitle
	Author      user.UserId
	ProgramType ProgramType
	CreatedAt   domain.CreatedAt
	UpdatedAt   *domain.UpdatedAt
}

func ProgramDerivationsFrom(
	derivedFrom *ProgramId,
	deriving []ProgramId,
) programDerivations {
	return programDerivations{
		DerivedFrom: derivedFrom,
		Deriving:    deriving,
	}
}

type programDerivations struct {
	DerivedFrom *ProgramId
	Deriving    []ProgramId
}

func ProgramDescriptionFrom(
	imageSrcs ProgramImageSrcs,
	text ProgramText,
) programDescription {
	return programDescription{
		ImageSrcs: imageSrcs,
		Text:      text,
	}
}

type programDescription struct {
	ImageSrcs ProgramImageSrcs
	Text      ProgramText
}

type ProgramUpdates struct {
	Title     *ProgramTitle
	ImageSrcs *ProgramImageSrcs
	Text      *ProgramText
}

type ProgramUpdateTargets struct {
	Title     ProgramTitle
	ImageSrcs ProgramImageSrcs
	Text      ProgramText
}
