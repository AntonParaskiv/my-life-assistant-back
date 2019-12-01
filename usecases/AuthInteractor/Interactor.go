package AuthInteractor

type Interactor struct {
	userRepository     UserRepositoryInterface
	sessionRepository  SessionRepositoryInterface
	sessionIdGenerator SessionIdGeneratorInterface
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetUserRepository(userRepository UserRepositoryInterface) *Interactor {
	i.userRepository = userRepository
	return i
}

func (i *Interactor) SetSessionRepository(sessionRepository SessionRepositoryInterface) *Interactor {
	i.sessionRepository = sessionRepository
	return i
}

func (i *Interactor) SetSessionIdGenerator(sessionIdGenerator SessionIdGeneratorInterface) *Interactor {
	i.sessionIdGenerator = sessionIdGenerator
	return i
}
