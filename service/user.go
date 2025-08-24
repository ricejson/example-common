package service

import (
	"github.com/ricejson/example-common/model"
)

type UserService interface {
	getUser(user model.User) model.User
}
