package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

const (
	NAME_MAX_LENGTH   = domain.ACT_NAME_MAX_LENGTH
	IMAGES_MAX_NUMBER = domain.ACT_IMAGES_MAX_NUMBER
	TEXT_MAX_LENGTH   = domain.ACT_TEXT_MAX_LENGTH
)

type ActId domain.Id
type ActCode string
type ActName string
type ActVersion uint

type ActImageSrcs []string
type ActText string

type Act struct {
	id   *ActId
	code ActCode

	actType ActType
	name    ActName
	version ActVersion
	author  user.UserId

	createdAt domain.CreatedAt
	updatedAt *domain.UpdatedAt

	imageSrcs       ActImageSrcs
	text            ActText
	characteristics ActCharacteristics
}

func (a *Act) Update(updater user.User, updates ActUpdates) error {
	if a.Base().Author != updater.Id() {
		return domain.ErrUnauthorized
	}
	return nil
}

func (a *Act) UpdateTargets() ActUpdateTargets {
	return ActUpdateTargets{
		ImageSrcs:       a.imageSrcs,
		Text:            a.text,
		Characteristics: a.characteristics,
	}
}

func (a *Act) IsPersisted() bool {
	return a.Id() != nil
}

func (a *Act) Id() *ActId {
	return a.id
}

func (a *Act) Code() ActCode {
	return a.code
}

func (a *Act) Base() ActBase {
	return ActBase{
		ActType:    a.actType,
		ActName:    a.name,
		ActVersion: a.version,
		Author:     a.author,
	}
}

func (a *Act) ActMetadata() ActMetadata {
	return ActMetadata{
		CreatedAt: a.createdAt,
		UpdatedAt: a.updatedAt,
	}
}

func (a *Act) ActDescription() ActDescription {
	return ActDescription{
		ImageSrcs:       a.imageSrcs,
		Text:            a.text,
		Characteristics: a.characteristics,
	}
}
