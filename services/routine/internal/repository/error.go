package repository

type RepositoryError struct {
	message string
}

func (e RepositoryError) Error() string {
	return e.message
}

var ErrNoTransaction = RepositoryError{"no transaction"}

var ErrNotFound = RepositoryError{"not found"}
