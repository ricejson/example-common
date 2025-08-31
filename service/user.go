package service

import (
	"github.com/ricejson/example-common/model"
)

type UserService interface {
	GetUser() (model.User, error)
}
