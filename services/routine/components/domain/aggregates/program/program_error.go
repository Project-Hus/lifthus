package program

type ProgramError struct {
	message string
}

func (e ProgramError) Error() string {
	return e.message
}

var ErrInvalidContents = ProgramError{"invalid contents"}
