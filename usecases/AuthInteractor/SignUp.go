package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/pkg/errors"
)

func (i *Interactor) SignUp(user *User.User) (err error) {
	isExist, err := i.IsUserExist(user)
	if err != nil {
		err = errors.Errorf("check if user exist failed")
		return
	}
	if isExist {
		err = errors.Errorf("user already exist")
		return
	}

	err = i.addUser(user)
	if err != nil {
		err = errors.Errorf("add user failed: %s", err.Error())
		return
	}

	return
}

func (i *Interactor) IsUserExist(user *User.User) (isExist bool, err error) {
	isExist, err = i.userRepository.IsUserExist(user)
	return
}

func (i *Interactor) addUser(user *User.User) (err error) {
	err = i.userRepository.AddUser(user)
	return
}
