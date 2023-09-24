package program

func DerivedFrom(
	program *Program,
	version uint,
) ProgramDerivedFrom {
	return ProgramDerivedFrom{
		program: program,
		version: version,
	}
}

type ProgramDerivedFrom struct {
	program *Program
	version uint
}
