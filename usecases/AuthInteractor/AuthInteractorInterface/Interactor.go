package AuthInteractorInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGeneratorInterface"
)

type Interactor interface {
	SetUserRepository(userRepository UserRepositoryInterface.Repository) Interactor
	SetSessionRepository(sessionRepository SessionRepositoryInterface.Repository) Interactor
	SetSessionIdGenerator(sessionIdGenerator SessionIdGeneratorInterface.Generator) Interactor
	IsUserExist(user UserInterface.User) (isExist bool, err error)
	IsUserValid(user UserInterface.User) (isValid bool, err error)
	SignIn(user UserInterface.User) (session SessionInterface.Session, err error)
	SignUp(user UserInterface.User) (err error)
}
