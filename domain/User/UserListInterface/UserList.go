package UserListInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
)

type UserList interface {
	AddUser(user UserInterface.User) (err error)
	GetUserByEmail(email string) (user UserInterface.User)
	IsUserExist(user UserInterface.User) (isExist bool)
}
