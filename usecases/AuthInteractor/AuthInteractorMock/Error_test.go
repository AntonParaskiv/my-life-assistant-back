package AuthInteractorMock

import (
	"reflect"
	"testing"
)

func TestInteractor_SimulateError(t *testing.T) {
	type fields struct {
		simulateErrorStepMatch int
		simulateErrorFlag      bool
	}
	type args struct {
		stepMatch int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				simulateErrorStepMatch: 0,
				simulateErrorFlag:      false,
			},
			args: args{
				stepMatch: 3,
			},
			want: &Interactor{
				simulateErrorStepMatch: 3,
				simulateErrorFlag:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				simulateErrorStepMatch: tt.fields.simulateErrorStepMatch,
				simulateErrorFlag:      tt.fields.simulateErrorFlag,
			}
			if got := i.SimulateError(tt.args.stepMatch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulateError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_IsSetSimulateError(t *testing.T) {
	type fields struct {
		simulateErrorStepMatch int
		simulateErrorFlag      bool
	}
	tests := []struct {
		name                   string
		fields                 fields
		wantIsSetSimulateError bool
		wantInteractor         *Interactor
	}{
		{
			name: "True",
			fields: fields{
				simulateErrorFlag: true,
			},
			wantIsSetSimulateError: true,
			wantInteractor: &Interactor{
				simulateErrorStepMatch: 0,
				simulateErrorFlag:      true,
			},
		},
		{
			name: "False with step",
			fields: fields{
				simulateErrorStepMatch: 3,
				simulateErrorFlag:      true,
			},
			wantIsSetSimulateError: false,
			wantInteractor: &Interactor{
				simulateErrorStepMatch: 2,
				simulateErrorFlag:      true,
			},
		},
		{
			name: "False",
			fields: fields{
				simulateErrorStepMatch: 0,
				simulateErrorFlag:      false,
			},
			wantIsSetSimulateError: false,
			wantInteractor: &Interactor{
				simulateErrorStepMatch: 0,
				simulateErrorFlag:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				simulateErrorStepMatch: tt.fields.simulateErrorStepMatch,
				simulateErrorFlag:      tt.fields.simulateErrorFlag,
			}
			if gotIsSetSimulateError := i.IsSetSimulateError(); gotIsSetSimulateError != tt.wantIsSetSimulateError {
				t.Errorf("IsSetSimulateError() = %v, want %v", gotIsSetSimulateError, tt.wantIsSetSimulateError)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}

func TestInteractor_Error(t *testing.T) {
	type fields struct {
		simulateErrorFlag bool
	}
	tests := []struct {
		name           string
		fields         fields
		wantErr        bool
		wantInteractor *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				simulateErrorFlag: true,
			},
			wantErr: true,
			wantInteractor: &Interactor{
				simulateErrorFlag: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := i.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}
