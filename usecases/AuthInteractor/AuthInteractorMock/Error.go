package AuthInteractorMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

func (i *Interactor) SimulateError(stepMatch int) *Interactor {
	i.simulateErrorStepMatch = stepMatch
	i.simulateErrorFlag = true
	return i
}

func (i *Interactor) IsSetSimulateError() (isSetSimulateError bool) {
	if !i.simulateErrorFlag {
		return
	}

	if i.simulateErrorStepMatch > 0 {
		i.simulateErrorStepMatch--
		return
	}

	isSetSimulateError = true
	return
}

func (i *Interactor) Error() (err error) {
	i.simulateErrorFlag = false
	err = errors.Errorf(ErrorSimulated)
	return
}
