package UserRepositoryMemory

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserListInterface"
	"github.com/pkg/errors"
)

type Repository struct {
	userList UserListInterface.UserListInterface
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetUserList(userList UserListInterface.UserListInterface) *Repository {
	r.userList = userList
	return r
}

func (r *Repository) AddUser(user UserInterface.User) (err error) {
	isExist, _ := r.IsUserExist(user)
	if isExist {
		err = errors.Errorf("user already exist")
		return
	}

	r.addUser(user)
	return
}

func (r *Repository) Auth(user UserInterface.User) (isValid bool, err error) {
	listUser := r.userList.GetUserByEmail(user.Email())
	if listUser == nil {
		err = errors.Errorf("user doesn't exist")
		return
	}

	if listUser.Password() != user.Password() {
		err = errors.Errorf("password doesn't match")
		return
	}

	isValid = true
	return
}

func (r *Repository) IsUserExist(user UserInterface.User) (isExist bool, err error) {
	isExist = r.userList.IsUserExist(user)
	return
}

func (r *Repository) addUser(user UserInterface.User) *Repository {
	_ = r.userList.AddUser(user)
	return r
}
