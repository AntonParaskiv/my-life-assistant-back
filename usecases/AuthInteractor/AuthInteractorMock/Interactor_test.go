package AuthInteractorMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/AuthInteractor/AuthInteractorInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGeneratorInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGeneratorMock"
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

func TestInteractor_SetIsUserExist(t *testing.T) {
	type fields struct {
		isUserExist bool
	}
	type args struct {
		isUserExist bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name: "True",
			fields: fields{
				isUserExist: false,
			},
			args: args{
				isUserExist: true,
			},
			want: &Interactor{
				isUserExist: true,
			},
		},
		{
			name: "False",
			fields: fields{
				isUserExist: true,
			},
			args: args{
				isUserExist: false,
			},
			want: &Interactor{
				isUserExist: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				isUserExist: tt.fields.isUserExist,
			}
			if got := i.SetIsUserExist(tt.args.isUserExist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsUserExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetIsUserValid(t *testing.T) {
	type fields struct {
		isUserValid bool
	}
	type args struct {
		isUserValid bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name: "True",
			fields: fields{
				isUserValid: false,
			},
			args: args{
				isUserValid: true,
			},
			want: &Interactor{
				isUserValid: true,
			},
		},
		{
			name: "False",
			fields: fields{
				isUserValid: true,
			},
			args: args{
				isUserValid: false,
			},
			want: &Interactor{
				isUserValid: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				isUserValid: tt.fields.isUserValid,
			}
			if got := i.SetIsUserValid(tt.args.isUserValid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsUserValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetSession(t *testing.T) {
	type fields struct {
		session SessionInterface.Session
	}
	type args struct {
		session SessionInterface.Session
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
				session: Session.New().SetId("mySession"),
			},
			want: &Interactor{
				session: Session.New().SetId("mySession"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				session: tt.fields.session,
			}
			if got := i.SetSession(tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetUser(t *testing.T) {
	type fields struct {
		user UserInterface.User
	}
	type args struct {
		user UserInterface.User
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
				user: User.New().SetEmail("myUser"),
			},
			want: &Interactor{
				user: User.New().SetEmail("myUser"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				user: tt.fields.user,
			}
			if got := i.SetUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUser() = %v, want %v", got, tt.want)
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
		want   AuthInteractorInterface.Interactor
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
		want   AuthInteractorInterface.Interactor
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
		sessionIdGenerator SessionIdGeneratorInterface.Generator
	}
	type args struct {
		sessionIdGenerator SessionIdGeneratorInterface.Generator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   AuthInteractorInterface.Interactor
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			want: &Interactor{
				sessionIdGenerator: SessionIdGeneratorMock.New(),
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
