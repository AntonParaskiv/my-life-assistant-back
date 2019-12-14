package AuthInteractorMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
)

func (i *Interactor) SignIn(user UserInterface.User) (session SessionInterface.Session, err error) {
	i.SetUser(user)
	session = i.session

	if i.IsSetSimulateError() {
		err = i.Error()
		return
	}

	return
}

func (i *Interactor) IsUserValid(user UserInterface.User) (isValid bool, err error) {
	i.SetUser(user)
	isValid = i.isUserValid

	if i.IsSetSimulateError() {
		err = i.Error()
		return
	}

	return
}
