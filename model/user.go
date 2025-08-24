package model

type User struct {
	name string
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
