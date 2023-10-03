package repository

type RepositoryError struct {
	message string
}

func (e RepositoryError) Error() string {
	return e.message
}

var ErrNotFound = RepositoryError{"not found"}
