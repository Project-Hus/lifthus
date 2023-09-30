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

var ErrInvalidProgramVersions = ProgramError{"invalid program versions"}
var ErrInvalidDailyRoutines = ProgramError{"invalid daily routines"}
var ErrInvalidRoutineActs = ProgramError{"invalid routine acts"}
