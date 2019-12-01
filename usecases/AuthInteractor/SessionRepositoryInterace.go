package AuthInteractor

import "github.com/AntonParaskiv/my-life-assistant-back/domain/Session"

type SessionRepositoryInterface interface {
	IsSessionIdExist(session *Session.Session) (isExist bool)
	AddSession(session *Session.Session) (err error)
}
