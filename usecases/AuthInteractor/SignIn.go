package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User"
	"github.com/pkg/errors"
)

func (i *Interactor) SignIn(user *User.User) (sessionId string, err error) {
	// check auth
	err = i.userRepository.Auth(user)
	if err != nil {
		err = errors.Errorf("auth failed: %s", err.Error())
		return
	}

	// create session
	session := Session.New().SetUser(user)

	// generate session id and check if it exist
	for ok := true; ok; {
		i.sessionIdGenerator.Generate(session)
		ok = i.sessionRepository.IsSessionIdExist(session)
	}

	err = i.sessionRepository.AddSession(session)
	if err != nil {
		err = errors.Errorf("create session failed: %s", err.Error())
		return
	}

	sessionId = session.Id()
	return
}
