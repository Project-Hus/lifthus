package repostiroy

import "routine/internal/domain/aggregates/act"

func NewActRepository(actRepo actRepository) *ActRepository {
	return &ActRepository{actRepo}
}

type actRepository interface {
	save(act *act.Act) (*act.Act, error)
	findByCode(code act.ActCode) (*act.Act, error)
}

type ActRepository struct {
	actRepository
}

func (ar *ActRepository) Save(act *act.Act) (*act.Act, error) {
	return ar.actRepository.save(act)
}

func (ar *ActRepository) FindByCode(code act.ActCode) (*act.Act, error) {
	return ar.actRepository.findByCode(code)
}
