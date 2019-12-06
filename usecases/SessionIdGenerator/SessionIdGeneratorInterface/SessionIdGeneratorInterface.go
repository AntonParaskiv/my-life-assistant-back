package SessionIdGeneratorInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
)

type SessionIdGeneratorInterface interface {
	Generate(session *Session.Session) *Session.Session
}
