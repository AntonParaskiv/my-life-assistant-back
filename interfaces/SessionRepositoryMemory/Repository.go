package SessionRepositoryMemory

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionListInterface"
	"github.com/pkg/errors"
)

type Repository struct {
	sessionList       SessionListInterface.SessionList
	simulateErrorFlag bool
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetSessionList(sessionList SessionListInterface.SessionList) *Repository {
	r.sessionList = sessionList
	return r
}

func (r *Repository) AddSession(session SessionInterface.Session) (err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	isExist, _ := r.IsSessionIdExist(session)
	if isExist {
		err = errors.Errorf("session id already exist")
		return
	}

	r.addSession(session)
	return
}

func (r *Repository) IsSessionIdExist(session SessionInterface.Session) (isExist bool, err error) {
	isExist = r.sessionList.IsSessionIdExist(session)
	return
}

func (r *Repository) addSession(session SessionInterface.Session) *Repository {
	_ = r.sessionList.AddSession(session)
	return r
}
