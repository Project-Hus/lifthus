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

	currentTime := domain.CreatedAt(time.Now())

	return &Act{
		code:      ActCode(code),
		author:    author.Id(),
		actType:   actType,
		name:      name,
		text:      text,
		imageSrcs: imageSrcs,
		createdAt: currentTime,
		standard:  false,
	}, nil
}

func ActFrom(
	code ActCode,
	authorId user.UserId,
	actType ActType,
	name ActName,
	text ActText,
	imageSrcs ActImageSrcs,
	createdAt domain.CreatedAt,
	standard bool,
) *Act {
	return &Act{
		code:      code,
		author:    authorId,
		actType:   actType,
		name:      name,
		text:      text,
		imageSrcs: imageSrcs,
		createdAt: createdAt,
		standard:  standard,
	}
}
