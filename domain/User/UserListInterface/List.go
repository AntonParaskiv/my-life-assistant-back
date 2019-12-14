package UserListInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
)

type List interface {
	AddUser(user UserInterface.User) (err error)
	Add(user UserInterface.User) List
	GetUserByEmail(email string) (user UserInterface.User)
	IsUserExist(user UserInterface.User) (isExist bool)
}
