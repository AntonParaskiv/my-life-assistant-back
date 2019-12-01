package AuthInteractor

import "github.com/AntonParaskiv/my-life-assistant-back/domain/Session"

type SessionIdGeneratorInterface interface {
	Generate(session *Session.Session) *Session.Session
}
