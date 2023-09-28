package act

import (
	"routine/components/domain"
)

func ActVersionFrom(version ActVersionNumber, imageSrcs ActImageSrcs, text ActText, characteristics ActCharacteristics, createdAt domain.CreatedAt) *ActVersion {
	return &ActVersion{
		version:         version,
		imageSrcs:       imageSrcs,
		text:            text,
		characteristics: characteristics,
		createdAt:       createdAt,
	}
}
