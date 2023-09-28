package actversion

import (
	"routine/components/domain/aggregates/act"
	"routine/components/domain/aggregates/user"
)

func ActVersionBaseFrom(
	actType act.ActType,
	actName act.ActName,
	author user.UserId,
) ActVersionBase {
	return ActVersionBase{
		ActType: actType,
		Name:    actName,
		Author:  author,
	}
}

type ActVersionBase struct {
	ActType act.ActType
	Name    act.ActName
	Author  user.UserId
}

type ActVersionReference struct {
	code       act.ActCode
	version    act.ActVersion
	referenced bool
}
