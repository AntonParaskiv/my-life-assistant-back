package SessionRepositoryInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
)

type Repository interface {
	AddSession(session SessionInterface.Session) (err error)
	IsSessionIdExist(session SessionInterface.Session) (isExist bool, err error)
}
