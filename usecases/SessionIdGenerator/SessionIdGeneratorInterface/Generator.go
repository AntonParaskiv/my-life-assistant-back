package SessionIdGeneratorInterface

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
)

type Generator interface {
	Generate(session SessionInterface.Session) SessionInterface.Session
}
