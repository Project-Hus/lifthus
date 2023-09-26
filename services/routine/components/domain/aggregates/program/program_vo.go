package program

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

func NewProgramMetadata(
	author user.UserId,
	programType ProgramType,
	createdAt domain.CreatedAt,
) programMetadata {
	return ProgramMetadataFrom(author, programType, createdAt, nil)
}

func ProgramMetadataFrom(
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

func ProgramInfoFrom(
	title ProgramTitle,
	imageSrcs ProgramImageSrcs,
	description ProgramDescription,
) programInfo {
	return programInfo{
		Title:       title,
		ImageSrcs:   imageSrcs,
		Description: description,
	}
}

type programInfo struct {
	Title       ProgramTitle
	ImageSrcs   ProgramImageSrcs
	Description ProgramDescription
}
