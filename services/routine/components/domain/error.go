package domain

type DomainError struct {
	message string
}

func (e DomainError) Error() string {
	return e.message
}

var ErrUnauthorized = DomainError{"unauthorized"}
