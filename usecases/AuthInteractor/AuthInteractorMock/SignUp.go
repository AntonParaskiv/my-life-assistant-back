package AuthInteractorMock

import "github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"

func (i *Interactor) SignUp(user UserInterface.User) (err error) {
	i.SetUser(user)

	if i.IsSetSimulateError() {
		err = i.Error()
		return
	}

	return
}

func (i *Interactor) IsUserExist(user UserInterface.User) (isExist bool, err error) {
	i.SetUser(user)
	isExist = i.isUserExist

	if i.IsSetSimulateError() {
		err = i.Error()
		return
	}

	return
}
