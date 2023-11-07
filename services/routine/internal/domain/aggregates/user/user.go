package user

func UserFrom(id UserId) *User {
	return &User{id: id}
}

type UserId int64

type User struct {
	id UserId
}

func (u User) Id() UserId {
	return u.id
}
