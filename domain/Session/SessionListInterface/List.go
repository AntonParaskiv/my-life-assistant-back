package SessionListInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
)

type List interface {
	AddSession(session SessionInterface.Session) (err error)
	Add(session SessionInterface.Session) List
	GetSessionById(id string) (session SessionInterface.Session)
	IsSessionIdExist(session SessionInterface.Session) (isExist bool)
}
