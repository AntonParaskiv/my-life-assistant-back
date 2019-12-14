package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGeneratorInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGenerator/SessionIdGeneratorMock"
	"reflect"
	"testing"
)

func TestInteractor_generateUniqueSessionId(t *testing.T) {
	type fields struct {
		sessionRepository  SessionRepositoryInterface.Repository
		sessionIdGenerator SessionIdGeneratorInterface.Generator
	}
	type args struct {
		session SessionInterface.Session
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantSessionWithId SessionInterface.Session
		wantErr           bool
	}{
		{
			name: "Success",
			fields: fields{
				sessionRepository:  SessionRepositoryMock.New(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				session: Session.New().SetUser(User.New().SetEmail("my@example.com")),
			},
			wantSessionWithId: Session.New().
				SetUser(User.New().SetEmail("my@example.com")).
				SetId("dfaa445d48f45b55bd695d89e593063c"),
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				sessionRepository:  SessionRepositoryMock.New().SimulateError(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				session: Session.New().SetUser(User.New().SetEmail("my@example.com")),
			},
			wantSessionWithId: nil,
			wantErr:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				sessionRepository:  tt.fields.sessionRepository,
				sessionIdGenerator: tt.fields.sessionIdGenerator,
			}
			gotSessionWithId, err := i.generateUniqueSessionId(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateUniqueSessionId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSessionWithId, tt.wantSessionWithId) {
				t.Errorf("generateUniqueSessionId() gotSessionWithId = %v, want %v", gotSessionWithId, tt.wantSessionWithId)
			}
		})
	}
}

func TestInteractor_createSession(t *testing.T) {
	type fields struct {
		sessionRepository  SessionRepositoryInterface.Repository
		sessionIdGenerator SessionIdGeneratorInterface.Generator
	}
	type args struct {
		user UserInterface.User
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		wantSession           SessionInterface.Session
		wantErr               bool
		wantSessionRepository SessionRepositoryInterface.Repository
	}{
		{
			name: "Success",
			fields: fields{
				sessionRepository:  SessionRepositoryMock.New(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com"),
			},
			wantSession: Session.New().
				SetUser(User.New().SetEmail("my@example.com")).
				SetId("dfaa445d48f45b55bd695d89e593063c"),
			wantErr: false,
			wantSessionRepository: SessionRepositoryMock.New().SetSession(
				Session.New().
					SetUser(User.New().SetEmail("my@example.com")).
					SetId("dfaa445d48f45b55bd695d89e593063c")),
		},
		{
			name: "Error",
			fields: fields{
				sessionRepository:  SessionRepositoryMock.New().SimulateError(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com"),
			},
			wantSession:           nil,
			wantErr:               true,
			wantSessionRepository: SessionRepositoryMock.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				sessionRepository:  tt.fields.sessionRepository,
				sessionIdGenerator: tt.fields.sessionIdGenerator,
			}
			gotSession, err := i.createSession(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("createSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSession, tt.wantSession) {
				t.Errorf("createSession() gotSession = %v, want %v", gotSession, tt.wantSession)
			}
			if !reflect.DeepEqual(i.sessionRepository, tt.wantSessionRepository) {
				t.Errorf("sessionRepository = %v, want %v", i.sessionRepository, tt.wantSessionRepository)
			}
		})
	}
}

func TestInteractor_IsUserValid(t *testing.T) {
	type fields struct {
		userRepository     UserRepositoryInterface.Repository
		sessionRepository  SessionRepositoryInterface.Repository
		sessionIdGenerator SessionIdGeneratorInterface.Generator
	}
	type args struct {
		user UserInterface.User
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantIsValid bool
		wantErr     bool
	}{
		{
			name: "Success",
			fields: fields{
				userRepository: UserRepositoryMock.New().
					SetUser(
						User.New().SetEmail("my@example.com").SetPassword("myPassword"),
					),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantIsValid: true,
			wantErr:     false,
		},
		{
			name: "Auth Failed: Password mismatch",
			fields: fields{
				userRepository: UserRepositoryMock.New().
					SetUser(
						User.New().SetEmail("my@example.com").SetPassword("myPassword"),
					),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("anotherPassword"),
			},
			wantIsValid: false,
			wantErr:     false,
		},
		{
			name: "Auth Failed: User unknown",
			fields: fields{
				userRepository: UserRepositoryMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("anotherPassword"),
			},
			wantIsValid: false,
			wantErr:     false,
		},
		{
			name: "Error",
			fields: fields{
				userRepository: UserRepositoryMock.New().SimulateError(0).
					SetUser(
						User.New().SetEmail("my@example.com").SetPassword("myPassword"),
					),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantIsValid: false,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				userRepository:     tt.fields.userRepository,
				sessionRepository:  tt.fields.sessionRepository,
				sessionIdGenerator: tt.fields.sessionIdGenerator,
			}
			gotIsValid, err := i.IsUserValid(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUserValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("IsUserValid() gotIsValid = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}

func TestInteractor_SignIn(t *testing.T) {
	type fields struct {
		userRepository     UserRepositoryInterface.Repository
		sessionRepository  SessionRepositoryInterface.Repository
		sessionIdGenerator SessionIdGeneratorInterface.Generator
	}
	type args struct {
		user UserInterface.User
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		wantSession           SessionInterface.Session
		wantErr               bool
		wantSessionRepository SessionRepositoryInterface.Repository
	}{
		{
			name: "Success",
			fields: fields{
				userRepository:     UserRepositoryMock.New().SetUser(User.New().SetEmail("my@example.com").SetPassword("myPassword")),
				sessionRepository:  SessionRepositoryMock.New(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantSession: Session.New().
				SetUser(User.New().SetEmail("my@example.com").SetPassword("myPassword")).
				SetId("dfaa445d48f45b55bd695d89e593063c"),
			wantErr: false,
			wantSessionRepository: SessionRepositoryMock.New().SetSession(
				Session.New().
					SetUser(User.New().SetEmail("my@example.com").SetPassword("myPassword")).
					SetId("dfaa445d48f45b55bd695d89e593063c")),
		},
		{
			name: "Error check valid",
			fields: fields{
				userRepository:     UserRepositoryMock.New().SimulateError(0),
				sessionRepository:  SessionRepositoryMock.New(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantSession:           nil,
			wantErr:               true,
			wantSessionRepository: SessionRepositoryMock.New(),
		},
		{
			name: "Error not valid",
			fields: fields{
				userRepository:     UserRepositoryMock.New().SetUser(User.New().SetEmail("my@example.com").SetPassword("myPassword")),
				sessionRepository:  SessionRepositoryMock.New(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("anotherPassword"),
			},
			wantSession:           nil,
			wantErr:               true,
			wantSessionRepository: SessionRepositoryMock.New(),
		},
		{
			name: "Error create session",
			fields: fields{
				userRepository:     UserRepositoryMock.New().SetUser(User.New().SetEmail("my@example.com").SetPassword("myPassword")),
				sessionRepository:  SessionRepositoryMock.New().SimulateError(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantSession:           nil,
			wantErr:               true,
			wantSessionRepository: SessionRepositoryMock.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				userRepository:     tt.fields.userRepository,
				sessionRepository:  tt.fields.sessionRepository,
				sessionIdGenerator: tt.fields.sessionIdGenerator,
			}
			gotSession, err := i.SignIn(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSession, tt.wantSession) {
				t.Errorf("SignIn() gotSession = %v, want %v", gotSession, tt.wantSession)
			}
			if !reflect.DeepEqual(i.sessionRepository, tt.wantSessionRepository) {
				t.Errorf("sessionRepository = %v, want %v", i.sessionRepository, tt.wantSessionRepository)
			}
		})
	}
}
