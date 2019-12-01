package UserRepositoryMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User"
	"github.com/pkg/errors"
)

type Repository struct {
	userList UserListInterface
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetUserList(userList UserListInterface) *Repository {
	r.userList = userList
	return r
}

func (r *Repository) AddUser(user *User.User) (err error) {
	if r.IsUserExist(user) {
		err = errors.Errorf("user already exist")
		return
	}

	r.addUser(user)
	return
}

func (r *Repository) Auth(user *User.User) (err error) {
	listUser := r.userList.GetUserByEmail(user.Email())
	if listUser == nil {
		err = errors.Errorf("user doesn't exist")
		return
	}

	if listUser.Password() != user.Password() {
		err = errors.Errorf("password doesn't match")
		return
	}

	return
}

func (r *Repository) IsUserExist(user *User.User) (isExist bool) {
	isExist = r.userList.IsUserExist(user)
	return
}

func (r *Repository) addUser(user *User.User) *Repository {
	r.userList.AddUser(user)
	return r
}
