package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"github.com/pkg/errors"
)

func (i *Interactor) SignIn(user UserInterface.User) (sessionId string, err error) {
	// check auth
	isValid, err := i.userRepository.Auth(user)
	if err != nil {
		err = errors.Errorf("auth error: %s", err.Error())
		return
	}
	if !isValid {
		err = errors.Errorf("auth failed: check login/password")
		return
	}

	// create session
	session := Session.New().SetUser(user)

	// generate session id and check if it exist
	for ok := true; ok; {
		i.sessionIdGenerator.Generate(session)
		ok, err = i.sessionRepository.IsSessionIdExist(session)
	}

	err = i.sessionRepository.AddSession(session)
	if err != nil {
		err = errors.Errorf("create session failed: %s", err.Error())
		return
	}

	sessionId = session.Id()
	return
}
