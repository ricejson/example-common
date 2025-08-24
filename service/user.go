package service

import (
	"github.com/ricejson/rice-rpc/example-common/model"
)

type UserService interface {
	getUser(user model.User) model.User
}
