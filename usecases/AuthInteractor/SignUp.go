package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User"
	"github.com/pkg/errors"
)

func (i *Interactor) SignUp(user *User.User) (sessionId string, err error) {
	err = i.userRepository.AddUser(user)
	if err != nil {
		err = errors.Errorf("add user failed: %s", err.Error())
		return
	}

	sessionId, err = i.SignIn(user)
	return
}
