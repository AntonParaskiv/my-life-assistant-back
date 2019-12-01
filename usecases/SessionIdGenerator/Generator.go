package SessionIdGenerator

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session"
	"github.com/google/uuid"
)

type Generator struct {
}

func New() (g *Generator) {
	g = new(Generator)
	return
}

func (g *Generator) Generate(session *Session.Session) *Session.Session {
	newId := g.generateNewId()
	session.SetId(newId)
	return session
}

func (g *Generator) generateNewId() (id string) {
	id = uuid.New().String()
	return
}
