package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/SessionList"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/UserList"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryMock"
	"github.com/AntonParaskiv/my-life-assistant-back/usecases/SessionIdGeneratorMock"
	"reflect"
	"testing"
)

func TestInteractor_SignIn(t *testing.T) {
	type fields struct {
		userRepository     UserRepositoryInterface
		sessionRepository  SessionRepositoryInterface
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
		wantSessionRepository SessionRepositoryInterface
	}{
		{
			name: "Success",
			fields: fields{
				userRepository: UserRepositoryMock.New().
					SetUserList(
						UserList.New().
							Add(User.New().SetEmail("first@user.com").SetPassword("firstPassword")).
							Add(User.New().SetEmail("second@user.com").SetPassword("secondPassword")).
							Add(User.New().SetEmail("third@user.com").SetPassword("thirdPassword")),
					),
				sessionRepository: SessionRepositoryMock.New().
					SetSessionList(
						SessionList.New().
							Add(Session.New().
								SetUser(User.New().SetEmail("first@user.com").SetPassword("firstPassword")).
								SetId("b56fc12a106757098d27713141fba845"),
							),
					),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("second@user.com").SetPassword("secondPassword"),
			},
			wantSessionId: "2e961b3a09933bf36f3143684af693ba",
			wantErr:       false,
			wantSessionRepository: SessionRepositoryMock.New().
				SetSessionList(
					SessionList.New().
						Add(Session.New().
							SetUser(User.New().SetEmail("first@user.com").SetPassword("firstPassword")).
							SetId("b56fc12a106757098d27713141fba845"),
						).
						Add(Session.New().
							SetUser(User.New().SetEmail("second@user.com").SetPassword("secondPassword")).
							SetId("2e961b3a09933bf36f3143684af693ba"),
						),
				),
		},
		{
			name: "Error Auth Failed",
			fields: fields{
				userRepository: UserRepositoryMock.New().
					SetUserList(
						UserList.New().
							Add(User.New().SetEmail("first@user.com").SetPassword("firstPassword")).
							Add(User.New().SetEmail("second@user.com").SetPassword("secondPassword")).
							Add(User.New().SetEmail("third@user.com").SetPassword("thirdPassword")),
					),
				sessionRepository:  SessionRepositoryMock.New().SetSessionList(SessionList.New()),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("fourth@user.com").SetPassword("fourthPassword"),
			},
			wantSessionId:         "",
			wantErr:               true,
			wantSessionRepository: SessionRepositoryMock.New().SetSessionList(SessionList.New()),
		},
		{
			name: "Error Add Session Failed",
			fields: fields{
				userRepository: UserRepositoryMock.New().
					SetUserList(
						UserList.New().
							Add(User.New().SetEmail("first@user.com").SetPassword("firstPassword")).
							Add(User.New().SetEmail("second@user.com").SetPassword("secondPassword")).
							Add(User.New().SetEmail("third@user.com").SetPassword("thirdPassword")),
					),
				sessionRepository: SessionRepositoryMock.New().
					SetSessionList(
						SessionList.New().
							Add(Session.New().
								SetUser(User.New().SetEmail("first@user.com").SetPassword("firstPassword")).
								SetId("b56fc12a106757098d27713141fba845"),
							),
					).
					SimulateError(),
				sessionIdGenerator: SessionIdGeneratorMock.New(),
			},
			args: args{
				user: User.New().SetEmail("first@user.com").SetPassword("firstPassword"),
			},
			wantSessionId: "",
			wantErr:       true,
			wantSessionRepository: SessionRepositoryMock.New().
				SetSessionList(
					SessionList.New().
						Add(Session.New().
							SetUser(User.New().SetEmail("first@user.com").SetPassword("firstPassword")).
							SetId("b56fc12a106757098d27713141fba845"),
						),
				),
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
