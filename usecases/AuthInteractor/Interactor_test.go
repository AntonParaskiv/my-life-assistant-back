package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGenerator"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGeneratorInterface"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantI *Interactor
	}{
		{
			name:  "Success",
			wantI: &Interactor{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestInteractor_SetUserRepository(t *testing.T) {
	type fields struct {
		userRepository UserRepositoryInterface.Repository
	}
	type args struct {
		userRepository UserRepositoryInterface.Repository
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				userRepository: UserRepositoryMock.New(),
			},
			want: &Interactor{
				userRepository: UserRepositoryMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				userRepository: tt.fields.userRepository,
			}
			if got := i.SetUserRepository(tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetSessionRepository(t *testing.T) {
	type fields struct {
		sessionRepository SessionRepositoryInterface.Repository
	}
	type args struct {
		sessionRepository SessionRepositoryInterface.Repository
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				sessionRepository: SessionRepositoryMock.New(),
			},
			want: &Interactor{
				sessionRepository: SessionRepositoryMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				sessionRepository: tt.fields.sessionRepository,
			}
			if got := i.SetSessionRepository(tt.args.sessionRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSessionRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetSessionIdGenerator(t *testing.T) {
	type fields struct {
		sessionIdGenerator SessionIdGeneratorInterface.SessionIdGeneratorInterface
	}
	type args struct {
		sessionIdGenerator SessionIdGeneratorInterface.SessionIdGeneratorInterface
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				sessionIdGenerator: SessionIdGenerator.New(),
			},
			want: &Interactor{
				sessionIdGenerator: SessionIdGenerator.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				sessionIdGenerator: tt.fields.sessionIdGenerator,
			}
			if got := i.SetSessionIdGenerator(tt.args.sessionIdGenerator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSessionIdGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}
