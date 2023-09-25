package programset

func CreateProgramset(programId uint64, edition Edition) *Programset {
	return &Programset{
		programId: programId,
		editions:  []Edition{edition},
	}
}

func ProgramsetFrom(programId uint64, editions []Edition) *Programset {
	return &Programset{
		programId: programId,
		editions:  editions,
	}
}

type Programset struct {
	programId uint64
	editions  []Edition
}
