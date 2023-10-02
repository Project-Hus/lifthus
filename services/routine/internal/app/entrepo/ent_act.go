package entrepo

import (
	"routine/internal/domain/aggregates/act"
	"routine/internal/ent"
	"routine/pkg/db"
)

func NewEntActRepository() *EntActRepository {
	return &EntActRepository{c: db.EntClient()}
}

type EntActRepository struct {
	c *ent.Client
}

func (repo *EntActRepository) Save(act *act.Act) (*act.Act, error) {
	return nil, nil
}

func (repo *EntActRepository) FindByCode(code act.ActCode) (*act.Act, error) {
	return nil, nil
}
