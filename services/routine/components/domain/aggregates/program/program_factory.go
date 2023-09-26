package program

import "routine/components/domain"

func CreateProgram(md programMetadata, drv programDerivations, inf programInfo, firstVersion Version) (*Program, error) {
	if !IsProgramInfoValid(inf) {
		return nil, ErrInvalidProgramInfo
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	return &Program{
		id:          nil,
		code:        ProgramCode(code),
		author:      md.Author,
		programType: md.ProgramType,
		createdAt:   md.CreatedAt,
		updatedAt:   md.UpdatedAt,
		derivedFrom: drv.DerivedFrom,
		deriving:    drv.Deriving,
		title:       inf.Title,
		imageSrcs:   inf.ImageSrcs,
		description: inf.Description,
		versions:    []*Version{&firstVersion},
	}, nil
}

func IsProgramInfoValid(info programInfo) bool {
	return IsTitleValid(info.Title) &&
		IsImageSrcsValid(info.ImageSrcs) &&
		IsDescriptionValid(info.Description)
}

func ProgramFrom(id ProgramId, code ProgramCode, md programMetadata, drv programDerivations, inf programInfo, versions []*Version) *Program {
	return &Program{
		id:          &id,
		code:        code,
		author:      md.Author,
		programType: md.ProgramType,
		createdAt:   md.CreatedAt,
		updatedAt:   md.UpdatedAt,
		derivedFrom: drv.DerivedFrom,
		deriving:    drv.Deriving,
		title:       inf.Title,
		imageSrcs:   inf.ImageSrcs,
		description: inf.Description,
		versions:    versions,
	}
}
