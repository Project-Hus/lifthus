package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
)

const (
	NAME_MIN_LENGTH   = domain.ACT_NAME_MIN_LENGTH
	NAME_MAX_LENGTH   = domain.ACT_NAME_MAX_LENGTH
	IMAGES_MAX_NUMBER = domain.ACT_IMAGES_MAX_NUMBER
	TEXT_MIN_LENGTH   = domain.ACT_TEXT_MIN_LENGTH
	TEXT_MAX_LENGTH   = domain.ACT_TEXT_MAX_LENGTH
)

type ActId domain.Id
type ActCode string
type ActName string
type ActVersion uint

type ActImageSrc string
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

	imageSrcs       []ActImageSrc
	text            ActText
	characteristics ActCharacteristics
}

func (a *Act) Update(updater user.User, updates ActUpdates) (*Act, error) {
	if a.Base().Author != updater.Id() {
		return nil, domain.ErrUnauthorized
	}
	if !IsActUpdatesValid(updates) {
		return nil, ErrInvalidActInfo
	}
	if updates.ImageSrcs != nil {
		a.imageSrcs = *updates.ImageSrcs
	}
	if updates.Text != nil {
		a.text = *updates.Text
	}
	if updates.Characteristics != nil {
		a.characteristics = *updates.Characteristics
	}
	return a, nil
}

func (a *Act) UpdateTargets() ActUpdateTargets {
	return ActUpdateTargets{
		ImageSrcs:       a.imageSrcs,
		Text:            a.text,
		Characteristics: a.characteristics,
	}
}

func (a *Act) Delete(deleter user.User) (*Act, error) {
	if a.Base().Author != deleter.Id() {
		return nil, domain.ErrUnauthorized
	}
	return a, nil
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

func (a *Act) Metadata() ActMetadata {
	return ActMetadata{
		CreatedAt: a.createdAt,
		UpdatedAt: a.updatedAt,
	}
}

func (a *Act) Description() ActDescription {
	return ActDescription{
		ImageSrcs:       a.imageSrcs,
		Text:            a.text,
		Characteristics: a.characteristics,
	}
}
