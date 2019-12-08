package SessionListInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
)

type SessionList interface {
	AddSession(session SessionInterface.Session) (err error)
	GetSessionById(id string) (session SessionInterface.Session)
	IsSessionIdExist(session SessionInterface.Session) (isExist bool)
}
