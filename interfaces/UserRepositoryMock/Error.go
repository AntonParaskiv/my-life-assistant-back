package UserRepositoryMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

func (r *Repository) SimulateError() *Repository {
	r.simulateErrorFlag = true
	return r
}

func (r *Repository) IsSetSimulateError() (IsSetSimulateError bool) {
	return r.simulateErrorFlag
}

func (r *Repository) Error() (err error) {
	r.simulateErrorFlag = false
	err = errors.Errorf(ErrorSimulated)
	return
}
