package SessionRepositoryMock

import "github.com/AntonParaskiv/my-life-assistant-back/domain/Session"

type SessionListInterface interface {
	AddSession(session *Session.Session) (err error)
	GetSessionById(id string) (session *Session.Session)
	IsSessionIdExist(session *Session.Session) (isExist bool)
}
