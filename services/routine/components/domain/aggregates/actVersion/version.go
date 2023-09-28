package actversion

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/act"
)

type Act struct {
	id        *act.ActId
	code      act.ActCode
	actType   act.ActType
	name      act.ActName
	createdAt domain.CreatedAt

	imageSrcs       act.ActImageSrcs
	text            act.ActText
	characteristics act.ActCharacteristics
}
