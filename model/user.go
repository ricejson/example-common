package model

type User struct {
	Name string
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}
