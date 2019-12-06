package SessionIdGeneratorMock

import (
	"crypto/md5"
	"fmt"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"io"
)

type Generator struct {
}

func New() (g *Generator) {
	g = new(Generator)
	return
}

func (g *Generator) Generate(session *Session.Session) *Session.Session {
	newId := g.GenerateIdFromString(session.User().Email() + session.Id())
	session.SetId(newId)
	return session
}

func (g *Generator) GenerateIdFromString(input string) (id string) {
	id = g.md5(input)
	return
}

func (g *Generator) md5(input string) (hash string) {
	h := md5.New()
	io.WriteString(h, input)
	hash = fmt.Sprintf("%x", h.Sum(nil))
	return
}
