package SessionRepositoryMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session"
	"github.com/pkg/errors"
)

type Repository struct {
	sessionList       SessionListInterface
	simulateErrorFlag bool
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetSessionList(sessionList SessionListInterface) *Repository {
	r.sessionList = sessionList
	return r
}

func (r *Repository) AddSession(session *Session.Session) (err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	if r.IsSessionIdExist(session) {
		err = errors.Errorf("session id already exist")
		return
	}

	r.addSession(session)
	return
}

func (r *Repository) IsSessionIdExist(session *Session.Session) (isExist bool) {
	isExist = r.sessionList.IsSessionIdExist(session)
	return
}

func (r *Repository) addSession(session *Session.Session) *Repository {
	r.sessionList.AddSession(session)
	return r
}
