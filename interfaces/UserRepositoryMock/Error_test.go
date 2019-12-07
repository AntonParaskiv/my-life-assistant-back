package UserRepositoryMock

import (
	"reflect"
	"testing"
)

func TestRepository_SimulateError(t *testing.T) {
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
		want   *Repository
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
			want: &Repository{
				simulateErrorStepMatch: 3,
				simulateErrorFlag:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				simulateErrorStepMatch: tt.fields.simulateErrorStepMatch,
				simulateErrorFlag:      tt.fields.simulateErrorFlag,
			}
			if got := r.SimulateError(tt.args.stepMatch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulateError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_IsSetSimulateError(t *testing.T) {
	type fields struct {
		simulateErrorStepMatch int
		simulateErrorFlag      bool
	}
	tests := []struct {
		name                   string
		fields                 fields
		wantIsSetSimulateError bool
		wantRepository         *Repository
	}{
		{
			name: "True",
			fields: fields{
				simulateErrorFlag: true,
			},
			wantIsSetSimulateError: true,
			wantRepository: &Repository{
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
			wantRepository: &Repository{
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
			wantRepository: &Repository{
				simulateErrorStepMatch: 0,
				simulateErrorFlag:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				simulateErrorStepMatch: tt.fields.simulateErrorStepMatch,
				simulateErrorFlag:      tt.fields.simulateErrorFlag,
			}
			if gotIsSetSimulateError := r.IsSetSimulateError(); gotIsSetSimulateError != tt.wantIsSetSimulateError {
				t.Errorf("IsSetSimulateError() = %v, want %v", gotIsSetSimulateError, tt.wantIsSetSimulateError)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("Repository = %v, want %v", r, tt.wantRepository)
			}
		})
	}
}

func TestRepository_Error(t *testing.T) {
	type fields struct {
		simulateErrorFlag bool
	}
	tests := []struct {
		name           string
		fields         fields
		wantErr        bool
		wantRepository *Repository
	}{
		{
			name: "Success",
			fields: fields{
				simulateErrorFlag: true,
			},
			wantErr: true,
			wantRepository: &Repository{
				simulateErrorFlag: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := r.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("mock = %v, want %v", r, tt.wantRepository)
			}
		})
	}
}
