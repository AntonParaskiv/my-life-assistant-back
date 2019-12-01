package AuthInteractor

import "github.com/AntonParaskiv/my-life-assistant-back/domain/User"

type UserRepositoryInterface interface {
	AddUser(user *User.User) (err error)
	Auth(user *User.User) (err error)
}
