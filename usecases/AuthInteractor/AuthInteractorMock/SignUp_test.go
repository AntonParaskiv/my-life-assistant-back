package AuthInteractorMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"reflect"
	"testing"
)

func TestInteractor_IsUserExist(t *testing.T) {
	type fields struct {
		user              UserInterface.User
		isUserExist       bool
		simulateErrorFlag bool
	}
	type args struct {
		user UserInterface.User
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantIsExist    bool
		wantErr        bool
		wantInteractor *Interactor
	}{
		{
			name: "True",
			fields: fields{
				isUserExist: true,
			},
			args: args{
				user: User.New().SetEmail("myEmail").SetPassword("myPassword"),
			},
			wantIsExist: true,
			wantErr:     false,
			wantInteractor: New().
				SetUser(
					User.New().SetEmail("myEmail").SetPassword("myPassword"),
				).
				SetIsUserExist(true),
		},
		{
			name: "False",
			fields: fields{
				isUserExist: false,
			},
			args: args{
				user: User.New().SetEmail("myEmail").SetPassword("myPassword"),
			},
			wantIsExist: false,
			wantErr:     false,
			wantInteractor: New().
				SetUser(
					User.New().SetEmail("myEmail").SetPassword("myPassword"),
				).
				SetIsUserExist(false),
		},
		{
			name: "Error",
			fields: fields{
				isUserExist:       false,
				simulateErrorFlag: true,
			},
			args: args{
				user: User.New().SetEmail("myEmail").SetPassword("myPassword"),
			},
			wantIsExist: false,
			wantErr:     true,
			wantInteractor: New().
				SetUser(
					User.New().SetEmail("myEmail").SetPassword("myPassword"),
				).
				SetIsUserExist(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				user:              tt.fields.user,
				isUserExist:       tt.fields.isUserExist,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotIsExist, err := i.IsUserExist(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUserExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsExist != tt.wantIsExist {
				t.Errorf("IsUserExist() gotIsExist = %v, want %v", gotIsExist, tt.wantIsExist)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}

func TestInteractor_SignUp(t *testing.T) {
	type fields struct {
		user              UserInterface.User
		simulateErrorFlag bool
	}
	type args struct {
		user UserInterface.User
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantInteractor *Interactor
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				user: User.New().SetEmail("myEmail").SetPassword("myPassword"),
			},
			wantErr:        false,
			wantInteractor: New().SetUser(User.New().SetEmail("myEmail").SetPassword("myPassword")),
		},
		{
			name: "Error",
			fields: fields{
				simulateErrorFlag: true,
			},
			args: args{
				user: User.New().SetEmail("myEmail").SetPassword("myPassword"),
			},
			wantErr:        true,
			wantInteractor: New().SetUser(User.New().SetEmail("myEmail").SetPassword("myPassword")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				user:              tt.fields.user,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := i.SignUp(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}
