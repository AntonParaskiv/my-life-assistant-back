package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/AuthInteractor/AuthInteractorInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGeneratorInterface"
)

type Interactor struct {
	userRepository     UserRepositoryInterface.Repository
	sessionRepository  SessionRepositoryInterface.Repository
	sessionIdGenerator SessionIdGeneratorInterface.Generator
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
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
