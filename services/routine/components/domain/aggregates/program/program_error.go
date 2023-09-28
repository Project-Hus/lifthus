package program

type ProgramError struct {
	message string
}

func (e ProgramError) Error() string {
	return e.message
}

var ErrInvalidProgramInfo = ProgramError{"invalid program information"}
var ErrInvalidProgramVersions = ProgramError{"invalid program versions"}
