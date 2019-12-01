package UserRepositoryMock

import "github.com/AntonParaskiv/my-life-assistant-back/domain/User"

type UserListInterface interface {
	AddUser(user *User.User) (err error)
	GetUserByEmail(email string) (user *User.User)
	IsUserExist(user *User.User) (isExist bool)
}
