package act

import (
	"routine/internal/domain"
	"routine/internal/domain/aggregates/user"
)

type ActCode domain.Code
type ActName string
type ActText string
type ActImageSrcs []string

type Act struct {
	code      ActCode
	author    user.UserId
	actType   ActType
	name      ActName
	text      ActText
	imageSrcs ActImageSrcs
	createdAt domain.CreatedAt
	standard  bool
}

type ActUpdateTargets struct {
	ImageSrcs *ActImageSrcs
	Text      *ActText
}

func (a *Act) Update(author user.User, targets ActUpdateTargets) (*Act, error) {
	if a.Author() != author.Id() {
		return nil, domain.ErrUnauthorized
	}
	if !targets.IsValid() {
		return nil, ErrInvalidActInfo
	}

	if targets.ImageSrcs != nil {
		a.imageSrcs = *targets.ImageSrcs
	}
	if targets.Text != nil {
		a.text = *targets.Text
	}
	return a, nil
}

func (a Act) Code() ActCode {
	return a.code
}

func (a Act) Author() user.UserId {
	return a.author
}

func (a Act) Type() ActType {
	return a.actType
}

func (a Act) Name() ActName {
	return a.name
}

func (a Act) Text() ActText {
	return a.text
}

func (a Act) ImageSrcs() ActImageSrcs {
	return a.imageSrcs
}

func (a Act) CreatedAt() domain.CreatedAt {
	return a.createdAt
}

func (a Act) IsStandard() bool {
	return a.standard
}
