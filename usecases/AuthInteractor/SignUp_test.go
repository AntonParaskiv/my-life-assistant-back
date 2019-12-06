package AuthInteractor

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/UserRepositoryMock"
	"reflect"
	"testing"
)

func TestInteractor_addUser(t *testing.T) {
	type fields struct {
		userRepository UserRepositoryInterface.Repository
	}
	type args struct {
		user *User.User
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantUserRepository UserRepositoryInterface.Repository
	}{
		{
			name: "Success",
			fields: fields{
				userRepository: UserRepositoryMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantErr: false,
			wantUserRepository: UserRepositoryMock.New().SetUser(
				User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			),
		},
		{
			name: "Error",
			fields: fields{
				userRepository: UserRepositoryMock.New().SimulateError(0),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantErr:            true,
			wantUserRepository: UserRepositoryMock.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				userRepository: tt.fields.userRepository,
			}
			if err := i.addUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("addUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(i.userRepository, tt.wantUserRepository) {
				t.Errorf("userRepository = %v, want %v", i.userRepository, tt.wantUserRepository)
			}
		})
	}
}

func TestInteractor_IsUserExist(t *testing.T) {
	type fields struct {
		userRepository UserRepositoryInterface.Repository
	}
	type args struct {
		user *User.User
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantIsExist bool
		wantErr     bool
	}{
		{
			name: "Exist",
			fields: fields{
				userRepository: UserRepositoryMock.New().
					SetUser(
						User.New().SetEmail("my@example.com"),
					),
			},
			args: args{
				user: User.New().SetEmail("my@example.com"),
			},
			wantIsExist: true,
			wantErr:     false,
		},
		{
			name: "Not exist",
			fields: fields{
				userRepository: UserRepositoryMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com"),
			},
			wantIsExist: false,
			wantErr:     false,
		},
		{
			name: "Error",
			fields: fields{
				userRepository: UserRepositoryMock.New().SimulateError(0).
					SetUser(
						User.New().SetEmail("my@example.com"),
					),
			},
			args: args{
				user: User.New().SetEmail("my@example.com"),
			},
			wantIsExist: false,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				userRepository: tt.fields.userRepository,
			}
			gotIsExist, err := i.IsUserExist(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUserExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsExist != tt.wantIsExist {
				t.Errorf("IsUserExist() gotIsExist = %v, want %v", gotIsExist, tt.wantIsExist)
			}
		})
	}
}

func TestInteractor_SignUp(t *testing.T) {
	type fields struct {
		userRepository UserRepositoryInterface.Repository
	}
	type args struct {
		user *User.User
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantUserRepository UserRepositoryInterface.Repository
	}{
		{
			name: "Success",
			fields: fields{
				userRepository: UserRepositoryMock.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantErr: false,
			wantUserRepository: UserRepositoryMock.New().
				SetUser(
					User.New().SetEmail("my@example.com").SetPassword("myPassword"),
				),
		},
		{
			name: "Error Check User Exist",
			fields: fields{
				userRepository: UserRepositoryMock.New().SimulateError(0).
					SetUser(
						User.New().SetEmail("my@example.com").SetPassword("myPassword"),
					),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantErr: true,
			wantUserRepository: UserRepositoryMock.New().
				SetUser(
					User.New().SetEmail("my@example.com").SetPassword("myPassword"),
				),
		},
		{
			name: "Error User Exist",
			fields: fields{
				userRepository: UserRepositoryMock.New().
					SetUser(
						User.New().SetEmail("my@example.com").SetPassword("myPassword"),
					),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantErr: true,
			wantUserRepository: UserRepositoryMock.New().
				SetUser(
					User.New().SetEmail("my@example.com").SetPassword("myPassword"),
				),
		},
		{
			name: "Error User Add",
			fields: fields{
				userRepository: UserRepositoryMock.New().SimulateError(1),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantErr:            true,
			wantUserRepository: UserRepositoryMock.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				userRepository: tt.fields.userRepository,
			}
			if err := i.SignUp(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(i.userRepository, tt.wantUserRepository) {
				t.Errorf("userRepository = %v, want %v", i.userRepository, tt.wantUserRepository)
			}
		})
	}
}
