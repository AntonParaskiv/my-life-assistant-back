package SessionRepositoryMemory

import (
	"reflect"
	"testing"
)

func TestRepository_SimulateError(t *testing.T) {
	type fields struct {
		simulateErrorFlag bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Repository
	}{
		{
			name: "Success",
			fields: fields{
				simulateErrorFlag: false,
			},
			want: &Repository{
				simulateErrorFlag: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Repository{
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if got := m.SimulateError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulateError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_IsSetSimulateError(t *testing.T) {
	type fields struct {
		simulateErrorFlag bool
	}
	tests := []struct {
		name                   string
		fields                 fields
		wantIsSetSimulateError bool
	}{
		{
			name: "True",
			fields: fields{
				simulateErrorFlag: true,
			},
			wantIsSetSimulateError: true,
		},
		{
			name: "False",
			fields: fields{
				simulateErrorFlag: false,
			},
			wantIsSetSimulateError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Repository{
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if gotIsSetSimulateError := m.IsSetSimulateError(); gotIsSetSimulateError != tt.wantIsSetSimulateError {
				t.Errorf("IsSetSimulateError() = %v, want %v", gotIsSetSimulateError, tt.wantIsSetSimulateError)
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
			m := &Repository{
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := m.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.wantRepository) {
				t.Errorf("mock = %v, want %v", m, tt.wantRepository)
			}
		})
	}
}
