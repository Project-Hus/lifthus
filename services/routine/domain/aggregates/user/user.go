package domain

var UserFactory = userFactory{}

type userFactory struct{}

func (uf userFactory) From(id uint64) User {
	return User{id: id}
}

type User struct {
	id uint64
}

func (u User) GetId() uint64 {
	return u.id
}
