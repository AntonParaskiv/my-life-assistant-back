package UserRepositoryMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

func (r *Repository) SimulateError(stepMatch int) *Repository {
	r.simulateErrorStepMatch = stepMatch
	r.simulateErrorFlag = true
	return r
}

func (r *Repository) IsSetSimulateError() (isSetSimulateError bool) {
	if !r.simulateErrorFlag {
		return
	}

	if r.simulateErrorStepMatch > 0 {
		r.simulateErrorStepMatch--
		return
	}

	isSetSimulateError = true
	return
}

func (r *Repository) Error() (err error) {
	r.simulateErrorFlag = false
	err = errors.Errorf(ErrorSimulated)
	return
}
