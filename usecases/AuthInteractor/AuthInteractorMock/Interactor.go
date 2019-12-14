package AuthInteractorMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/AuthInteractor/AuthInteractorInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGeneratorInterface"
)

type Interactor struct {
	userRepository     UserRepositoryInterface.Repository
	sessionRepository  SessionRepositoryInterface.Repository
	sessionIdGenerator SessionIdGeneratorInterface.Generator

	user        UserInterface.User
	session     SessionInterface.Session
	isUserExist bool
	isUserValid bool

	simulateErrorStepMatch int
	simulateErrorFlag      bool
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetUser(user UserInterface.User) *Interactor {
	i.user = user
	return i
}

func (i *Interactor) SetSession(session SessionInterface.Session) *Interactor {
	i.session = session
	return i
}

func (i *Interactor) SetIsUserExist(isUserExist bool) *Interactor {
	i.isUserExist = isUserExist
	return i
}

func (i *Interactor) SetIsUserValid(isUserValid bool) *Interactor {
	i.isUserValid = isUserValid
	return i
}

func (i *Interactor) SetUserRepository(userRepository UserRepositoryInterface.Repository) AuthInteractorInterface.Interactor {
	i.userRepository = userRepository
	return i
}

func (i *Interactor) SetSessionRepository(sessionRepository SessionRepositoryInterface.Repository) AuthInteractorInterface.Interactor {
	i.sessionRepository = sessionRepository
	return i
}

func (i *Interactor) SetSessionIdGenerator(sessionIdGenerator SessionIdGeneratorInterface.Generator) AuthInteractorInterface.Interactor {
	i.sessionIdGenerator = sessionIdGenerator
	return i
}
