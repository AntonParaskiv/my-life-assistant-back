package AuthInteractorMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"reflect"
	"testing"
)

func TestInteractor_IsUserValid(t *testing.T) {
	type fields struct {
		user              UserInterface.User
		isUserValid       bool
		simulateErrorFlag bool
	}
	type args struct {
		user UserInterface.User
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantIsValid    bool
		wantErr        bool
		wantInteractor *Interactor
	}{
		{
			name: "Valid",
			fields: fields{
				isUserValid: true,
			},
			args: args{
				user: User.New().SetEmail("myEmail").SetPassword("myPassword"),
			},
			wantIsValid: true,
			wantErr:     false,
			wantInteractor: New().
				SetUser(
					User.New().SetEmail("myEmail").SetPassword("myPassword"),
				).
				SetIsUserValid(true),
		},
		{
			name: "InValid",
			fields: fields{
				isUserValid: false,
			},
			args: args{
				user: User.New().SetEmail("myEmail").SetPassword("myPassword"),
			},
			wantIsValid: false,
			wantErr:     false,
			wantInteractor: New().
				SetUser(
					User.New().SetEmail("myEmail").SetPassword("myPassword"),
				).
				SetIsUserValid(false),
		},
		{
			name: "Error",
			fields: fields{
				isUserValid:       false,
				simulateErrorFlag: true,
			},
			args: args{
				user: User.New().SetEmail("myEmail").SetPassword("myPassword"),
			},
			wantIsValid: false,
			wantErr:     true,
			wantInteractor: New().
				SetUser(
					User.New().SetEmail("myEmail").SetPassword("myPassword"),
				).
				SetIsUserValid(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				user:              tt.fields.user,
				isUserValid:       tt.fields.isUserValid,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotIsValid, err := i.IsUserValid(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUserValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("IsUserValid() gotIsValid = %v, want %v", gotIsValid, tt.wantIsValid)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}

func TestInteractor_SignIn(t *testing.T) {
	type fields struct {
		user              UserInterface.User
		session           SessionInterface.Session
		simulateErrorFlag bool
	}
	type args struct {
		user UserInterface.User
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantSession    SessionInterface.Session
		wantErr        bool
		wantInteractor *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				session: Session.New().
					SetId("mySessionId").
					SetUser(User.New().
						SetEmail("myEmail").
						SetPassword("myPassword")),
			},
			args: args{
				user: User.New().
					SetEmail("myEmail").
					SetPassword("myPassword"),
			},
			wantSession: Session.New().
				SetId("mySessionId").
				SetUser(User.New().
					SetEmail("myEmail").
					SetPassword("myPassword")),
			wantErr: false,
			wantInteractor: New().
				SetUser(User.New().
					SetEmail("myEmail").
					SetPassword("myPassword")).
				SetSession(Session.New().
					SetId("mySessionId").
					SetUser(User.New().SetEmail("myEmail").SetPassword("myPassword")),
				),
		},
		{
			name: "Error",
			fields: fields{
				session:           nil,
				simulateErrorFlag: true,
			},
			args: args{
				user: User.New().
					SetEmail("myEmail").
					SetPassword("myPassword"),
			},
			wantSession: nil,
			wantErr:     true,
			wantInteractor: New().
				SetUser(User.New().
					SetEmail("myEmail").
					SetPassword("myPassword")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				user:              tt.fields.user,
				session:           tt.fields.session,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotSession, err := i.SignIn(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSession, tt.wantSession) {
				t.Errorf("SignIn() gotSession = %v, want %v", gotSession, tt.wantSession)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}
