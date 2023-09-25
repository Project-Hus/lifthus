package program

func ProgramDerivationsFrom(
	derivedFrom *ProgramId,
	deriving []ProgramId,
) programDerivations {
	return programDerivations{
		DerivedFrom: derivedFrom,
		Deriving:    deriving,
	}
}

type programDerivations struct {
	DerivedFrom *ProgramId
	Deriving    []ProgramId
}
