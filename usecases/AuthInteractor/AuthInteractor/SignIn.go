package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"github.com/pkg/errors"
)

func (i *Interactor) SignIn(user UserInterface.User) (session SessionInterface.Session, err error) {
	isValid, err := i.IsUserValid(user)
	if err != nil {
		return
	}
	if !isValid {
		err = errors.Errorf("auth failed: check login/password")
		return
	}

	session, err = i.createSession(user)
	if err != nil {
		err = errors.Errorf("create session failed: %s", err.Error())
		return
	}

	return
}

func (i *Interactor) IsUserValid(user UserInterface.User) (isValid bool, err error) {
	isValid, err = i.userRepository.Auth(user)
	if err != nil {
		err = errors.Errorf("auth error: %s", err.Error())
		return
	}
	return
}

func (i *Interactor) createSession(user UserInterface.User) (session SessionInterface.Session, err error) {
	// TODO: replace constructor with Factory !
	session = Session.New().SetUser(user)
	session, err = i.generateUniqueSessionId(session)
	if err != nil {
		return
	}

	err = i.sessionRepository.AddSession(session)
	return
}

func (i *Interactor) generateUniqueSessionId(session SessionInterface.Session) (sessionWithId SessionInterface.Session, err error) {
	for ok := true; ok; {
		i.sessionIdGenerator.Generate(session)
		ok, err = i.sessionRepository.IsSessionIdExist(session)
	}
	if err != nil {
		err = errors.Errorf("create session failed: %s", err.Error())
		return
	}
	sessionWithId = session
	return
}
