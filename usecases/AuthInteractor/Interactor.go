package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
)

type Interactor struct {
	userRepository     UserRepositoryInterface.Repository
	sessionRepository  SessionRepositoryInterface.Repository
	sessionIdGenerator SessionIdGeneratorInterface
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetUserRepository(userRepository UserRepositoryInterface.Repository) *Interactor {
	i.userRepository = userRepository
	return i
}

func (i *Interactor) SetSessionRepository(sessionRepository SessionRepositoryInterface.Repository) *Interactor {
	i.sessionRepository = sessionRepository
	return i
}

func (i *Interactor) SetSessionIdGenerator(sessionIdGenerator SessionIdGeneratorInterface) *Interactor {
	i.sessionIdGenerator = sessionIdGenerator
	return i
}
