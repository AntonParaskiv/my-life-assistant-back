package UserRepositoryMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
)

type Repository struct {
	user              UserInterface.User
	stepMatch         int
	simulateErrorFlag bool
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetUser(user UserInterface.User) *Repository {
	r.user = user
	return r
}

func (r *Repository) User() (user UserInterface.User) {
	user = r.user
	return
}

func (r *Repository) AddUser(user UserInterface.User) (err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	r.SetUser(user)
	return
}

func (r *Repository) IsUserExist(user UserInterface.User) (isExist bool, err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	if r.User() == nil {
		return
	}

	isExist = r.User().Email() == user.Email()
	return
}

func (r *Repository) Auth(user UserInterface.User) (isValid bool, err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	if r.User() == nil {
		return
	}

	isValid = r.User().Email() == user.Email() && r.User().Password() == user.Password()
	return
}
