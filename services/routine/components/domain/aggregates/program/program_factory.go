package program

import (
	"fmt"
)

func Create(info *NewProgramInfo) (*Program, error) {
	switch {
	case len(info.Title) > TITLE_MAX_LENGTH:
		fallthrough
	case len(info.ImageSrcs) > IMAGES_MAX_COUNT:
		fallthrough
	case len(info.Description) > DESCRIPTION_MAX_LENGTH:
		return nil, fmt.Errorf("invalid program info")
	}
	return &Program{
		parent:      info.Parent,
		title:       info.Title,
		author:      info.Author,
		programType: info.ProgramType,
		iteration:   info.Iteration,
		imageSrcs:   info.ImageSrcs,
		description: info.Description,
	}, nil
}

func From(info *ProgramInfo) *Program {
	return &Program{
		id:          info.Id,
		slug:        info.Slug,
		code:        info.Code,
		parent:      info.Parent,
		title:       info.Title,
		author:      info.Author,
		programType: info.ProgramType,
		iteration:   info.Iteration,
		imageSrcs:   info.ImageSrcs,
		description: info.Description,
		createdAt:   info.createdAt,
		updatedAt:   info.updatedAt,
	}
}
