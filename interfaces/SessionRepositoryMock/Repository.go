package SessionRepositoryMock

type Repository struct {
	session           SessionInterface
	simulateErrorFlag bool
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetSession(session SessionInterface) *Repository {
	r.session = session
	return r
}

func (r *Repository) Session() (session SessionInterface) {
	session = r.session
	return
}

func (r *Repository) AddSession(session SessionInterface) (err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	r.SetSession(session)
	return
}

func (r *Repository) IsSessionIdExist(session SessionInterface) (isExist bool, err error) {
	if r.IsSetSimulateError() {
		err = r.Error()
		return
	}

	isExist = r.Session().Id() == session.Id()
	return
}
