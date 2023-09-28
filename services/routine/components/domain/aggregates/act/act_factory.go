package act

import (
	"routine/components/domain"
	"routine/components/domain/aggregates/user"
	"time"
)

func CreateAct(
	actType ActType,
	name ActName,
	author user.User,
	imageSrcs ActImageSrcs,
	text ActText,
	characteristics ActCharacteristics,
) (*Act, error) {
	if !name.IsValid() || !imageSrcs.IsValid() || !text.IsValid() {
		return nil, ErrInvalidActInfo
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}

	currentTime := domain.CreatedAt(time.Now())

	newVersion := ActVersionFrom(
		1,
		imageSrcs,
		text,
		characteristics,
		currentTime,
	)

	return &Act{
		code:      ActCode(code),
		actType:   actType,
		name:      name,
		author:    author.Id(),
		createdAt: currentTime,
		versions:  []*ActVersion{newVersion},
	}, nil
}

func ActFrom(
	code ActCode,
	actType ActType,
	name ActName,
	author user.User,
	createAt domain.CreatedAt,
	versions ActVersions,
) (*Act, error) {
	if !versions.IsValid() {
		return nil, ErrUnsortedActVersions
	}
	return &Act{
		code: code,

		actType:   actType,
		name:      name,
		author:    author.Id(),
		createdAt: createAt,

		versions: versions,
	}, nil
}
