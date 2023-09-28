package act

import (
	"routine/components/domain"
)

func ActVersionFrom(code ActVersionCode, version ActVersionNumber, imageSrcs ActImageSrcs, text ActText, createdAt domain.CreatedAt) *ActVersion {
	return &ActVersion{
		code:      code,
		version:   version,
		imageSrcs: imageSrcs,
		text:      text,
		createdAt: createdAt,
	}
}
