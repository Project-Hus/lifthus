package actversion

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/act"
	"routine/components/domain/aggregates/user"
)

type ActVersion struct {
	code    act.ActCode
	version act.ActVersion

	actType act.ActType
	name    act.ActName
	author  user.UserId

	createdAt domain.CreatedAt

	imageSrcs       []act.ActImageSrc
	text            act.ActText
	characteristics act.ActCharacteristics
}

// TODO: deletion should only be available when there is no RoutineAct referencing corresponding ActVersion
func (av *ActVersion) Delete(deleter user.User) (*ActVersion, error) {
	return nil, nil
}

func (av *ActVersion) Code() act.ActCode {
	return av.code
}

func (av *ActVersion) Version() act.ActVersion {
	return av.version
}

func (av *ActVersion) Base() ActVersionBase {
	return ActVersionBase{
		ActType: av.actType,
		Name:    av.name,
		Author:  av.author,
	}
}

func (av *ActVersion) CreatedAt() domain.CreatedAt {
	return av.createdAt
}

func (av *ActVersion) Description() act.ActDescription {
	return act.ActDescription{
		ImageSrcs:       av.imageSrcs,
		Text:            av.text,
		Characteristics: av.characteristics,
	}
}
