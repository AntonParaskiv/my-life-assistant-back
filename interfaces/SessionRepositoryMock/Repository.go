package SessionRepositoryMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
)

type Repository struct {
	session           SessionInterface.Session
	simulateErrorFlag bool
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetSession(session SessionInterface.Session) SessionRepositoryInterface.Repository {
	r.session = session
	return r
}

func (r *Repository) Session() (session SessionInterface.Session) {
	session = r.session
	return
}

func (r *Repository) AddSession(session SessionInterface.Session) (err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	r.SetSession(session)
	return
}

func (r *Repository) IsSessionIdExist(session SessionInterface.Session) (isExist bool, err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	if r.Session() == nil {
		return
	}

	isExist = r.Session().Id() == session.Id()
	return
}
