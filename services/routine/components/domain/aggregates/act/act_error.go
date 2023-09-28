package act

type ActError struct {
	message string
}

func (e ActError) Error() string {
	return e.message
}

var ErrInvalidActInfo = ActError{"invalid act info"}
