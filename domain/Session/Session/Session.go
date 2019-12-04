package Session

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
)

type Session struct {
	user UserInterface.User
	id   string
}

func New() (s *Session) {
	s = new(Session)
	return
}

func (s *Session) SetUser(user UserInterface.User) *Session {
	s.user = user
	return s
}

func (s *Session) SetId(id string) *Session {
	s.id = id
	return s
}

func (s *Session) User() (user UserInterface.User) {
	user = s.user
	return
}

func (s *Session) Id() (id string) {
	id = s.id
	return
}
