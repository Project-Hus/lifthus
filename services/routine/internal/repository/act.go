package repository

import "routine/internal/domain/aggregates/act"

func NewActRepository(actRepo actRepository) *ActRepository {
	return &ActRepository{actRepo}
}

type actRepository interface {
	Save(act *act.Act) (*act.Act, error)
	FindByCode(code act.ActCode) (*act.Act, error)
}

type ActRepository struct {
	actRepository
}

func (ar *ActRepository) Save(act *act.Act) (*act.Act, error) {
	return ar.actRepository.Save(act)
}

func (ar *ActRepository) FindByCode(code act.ActCode) (*act.Act, error) {
	return ar.actRepository.FindByCode(code)
}
