package dto

type CreateActDto struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	Author      uint64  `json:"author,omitempty"`
	Image       *string `json:"image,omitempty"`
	Description *string `json:"description,omitempty"`

	// weight/cardio
	Weight     bool `json:"weight,omitempty"`
	Bodyweight bool `json:"bodyweight,omitempty"`
	Cardio     bool `json:"cardio,omitempty"`

	// upper/lower/full body
	Upper bool `json:"upper,omitempty"`
	Lower bool `json:"lower,omitempty"`
	Full  bool `json:"full,omitempty"`

	// prime movers
	Arms      bool `json:"arms,omitempty"`
	Shoulders bool `json:"shoulders,omitempty"`
	Chest     bool `json:"chest,omitempty"`
	Core      bool `json:"core,omitempty"`
	UpperBack bool `json:"upper_back,omitempty"`
	LowerBack bool `json:"lower_back,omitempty"`
	Glute     bool `json:"glute,omitempty"`
	LegsFront bool `json:"legs_front,omitempty"`
	LegsBack  bool `json:"legs_back,omitempty"`
	Etc       bool `json:"etc,omitempty"`
}
