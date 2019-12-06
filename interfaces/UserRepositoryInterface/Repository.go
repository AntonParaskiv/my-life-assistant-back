package UserRepositoryInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
)

type Repository interface {
	AddUser(user UserInterface.User) (err error)
	Auth(user UserInterface.User) (isValid bool, err error)
	IsUserExist(user UserInterface.User) (isExist bool, err error)
}
