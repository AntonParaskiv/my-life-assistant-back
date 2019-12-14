package SessionRepositoryMemory

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionListInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/pkg/errors"
)

type Repository struct {
	sessionList SessionListInterface.List
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetSessionList(sessionList SessionListInterface.List) SessionRepositoryInterface.Repository {
	r.sessionList = sessionList
	return r
}

func (r *Repository) AddSession(session SessionInterface.Session) (err error) {
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

func (r *Repository) addSession(session SessionInterface.Session) SessionRepositoryInterface.Repository {
	_ = r.sessionList.AddSession(session)
	return r
}
