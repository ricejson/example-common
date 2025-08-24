package model

type User struct {
	name string
}

func NewUser() *User {
	return &User{}
}

func (u *User) getName() string {
	return u.name
}

func (u *User) setName(name string) {
	u.name = name
}
