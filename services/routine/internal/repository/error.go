package repository

type RepositoryError struct {
	message string
}

func (e RepositoryError) Error() string {
	return e.message
}

var ErrNoTransaction = RepositoryError{"no transaction"}

func IsNotFound(err error) bool {
	if _, ok := err.(RepositoryError); ok && err.Error() == ErrNotFound.Error() {
		return true
	}
	return false
}

var ErrNotFound = RepositoryError{"not found"}

var ErrInvalidActType = RepositoryError{"invalid act type"}
