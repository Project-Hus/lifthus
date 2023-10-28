package program

type ProgramError struct {
	message string
}

func (e ProgramError) Error() string {
	return e.message
}

var ErrInvalidProgramInfo = ProgramError{"invalid program information"}
var ErrInvalidProgramTitle = ProgramError{"invalid program title"}
var ErrInvalidProgramImageSrcs = ProgramError{"invalid program image sources"}
var ErrInvalidProgramText = ProgramError{"invalid program text"}

var ErrInvalidProgramReleases = ProgramError{"invalid program releases"}
var ErrInvalidRoutines = ProgramError{"invalid routines"}
var ErrInvalidRoutineActs = ProgramError{"invalid routine acts"}
