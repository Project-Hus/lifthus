package act

import (
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
	"time"
)

func CreateAct(
	actType ActType,
	name ActName,
	author user.User,
	imageSrcs ActImageSrcs,
	text ActText,
) (*Act, error) {
	if !name.IsValid() || !imageSrcs.IsValid() || !text.IsValid() {
		return nil, ErrInvalidActInfo
	}
	code, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}
	vCode, err := domain.RandomHexCode()
	if err != nil {
		return nil, err
	}

	currentTime := domain.CreatedAt(time.Now())

	newVersion := ActVersionFrom(
		ActVersionCode(vCode),
		1,
		imageSrcs,
		text,
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
	authorId user.UserId,
	createAt domain.CreatedAt,
	versions ActVersions,
) (*Act, error) {
	if !versions.IsValid() {
		return nil, ErrInvalidActVersions
	}
	return &Act{
		code: code,

		actType:   actType,
		name:      name,
		author:    authorId,
		createdAt: createAt,

		versions: versions,
	}, nil
}
