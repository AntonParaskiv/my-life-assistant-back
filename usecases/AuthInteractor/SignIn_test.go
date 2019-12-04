package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGeneratorMock"
	"reflect"
	"testing"
)

func TestInteractor_SignIn(t *testing.T) {
	type fields struct {
		userRepository     UserRepositoryInterface.Repository
		sessionRepository  SessionRepositoryInterface.Repository
		sessionIdGenerator SessionIdGeneratorInterface
	}
	type args struct {
		user *User.User
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		wantSessionId         string
		wantErr               bool
		wantSessionRepository SessionRepositoryInterface.Repository
	}{
		{
			name: "Success",
			fields: fields{
				userRepository: UserRepositoryMock.New().SetUser(
					User.New().SetEmail("my@example.com").SetPassword("myPassword"),
				),
				sessionRepository:  SessionRepositoryMock.New(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantSessionId: "dfaa445d48f45b55bd695d89e593063c",
			wantErr:       false,
			wantSessionRepository: SessionRepositoryMock.New().SetSession(
				Session.New().
					SetUser(User.New().SetEmail("my@example.com").SetPassword("myPassword")).
					SetId("dfaa445d48f45b55bd695d89e593063c"),
			),
		},
		{
			name: "Auth Error",
			fields: fields{
				userRepository: UserRepositoryMock.New().SimulateError(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantSessionId:         "",
			wantErr:               true,
			wantSessionRepository: nil,
		},
		{
			name: "Auth Failed",
			fields: fields{
				userRepository: UserRepositoryMock.New().SetUser(
					User.New().SetEmail("my@example.com").SetPassword("myPassword"),
				),
			},
			args: args{
				User.New().SetEmail("my@example.com").SetPassword("anotherPassword"),
			},
			wantSessionId:         "",
			wantErr:               true,
			wantSessionRepository: nil,
		},
		{
			name: "Error Add Session Failed",
			fields: fields{
				userRepository: UserRepositoryMock.New().SetUser(
					User.New().SetEmail("my@example.com").SetPassword("myPassword"),
				),
				sessionRepository:  SessionRepositoryMock.New().SimulateError(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantSessionId:         "",
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
			gotSessionId, err := i.SignIn(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSessionId != tt.wantSessionId {
				t.Errorf("SignIn() gotSessionId = %v, want %v", gotSessionId, tt.wantSessionId)
			}
			if !reflect.DeepEqual(tt.fields.sessionRepository, tt.wantSessionRepository) {
				t.Errorf("sessionRepository = %v, want %v", tt.fields.sessionRepository, tt.wantSessionRepository)
			}
		})
	}
}
