package program

type ProgramError struct {
	message string
}

func (e ProgramError) Error() string {
	return e.message
}

var ErrInvalidProgramInfo = ProgramError{"invalid program information"}

var ErrExistingDerivingProgram = ProgramError{"existing deriving program"}
