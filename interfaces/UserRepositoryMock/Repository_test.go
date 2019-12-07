package UserRepositoryMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantR *Repository
	}{
		{
			name:  "Success",
			wantR: &Repository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := New(); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("New() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestRepository_SetUser(t *testing.T) {
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
		want   *Repository
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				user: User.New(),
			},
			want: &Repository{
				user: User.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				user: tt.fields.user,
			}
			if got := r.SetUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_User(t *testing.T) {
	type fields struct {
		user UserInterface.User
	}
	tests := []struct {
		name     string
		fields   fields
		wantUser UserInterface.User
	}{
		{
			name: "Success",
			fields: fields{
				user: User.New(),
			},
			wantUser: User.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				user: tt.fields.user,
			}
			if gotUser := r.User(); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("User() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestRepository_AddUser(t *testing.T) {
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
		wantRepository *Repository
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				user: User.New().SetEmail("my@example.com"),
			},
			wantErr: false,
			wantRepository: &Repository{
				user: User.New().SetEmail("my@example.com"),
			},
		},
		{
			name: "Error",
			fields: fields{
				simulateErrorFlag: true,
			},
			args: args{
				user: User.New().SetEmail("my@example.com"),
			},
			wantErr: true,
			wantRepository: &Repository{
				user: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				user:              tt.fields.user,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := r.AddUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("Repository = %v, wantErr %v", r, tt.wantRepository)
			}
		})
	}
}

func TestRepository_IsUserExist(t *testing.T) {
	type fields struct {
		user              UserInterface.User
		simulateErrorFlag bool
	}
	type args struct {
		user UserInterface.User
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantIsExist bool
		wantErr     bool
	}{
		{
			name: "True",
			fields: fields{
				user: User.New().SetEmail("my@example.com"),
			},
			args: args{
				user: User.New().SetEmail("my@example.com"),
			},
			wantIsExist: true,
			wantErr:     false,
		},
		{
			name: "True",
			fields: fields{
				user: User.New().SetEmail("my@example.com"),
			},
			args: args{
				user: User.New().SetEmail("another@example.com"),
			},
			wantIsExist: false,
			wantErr:     false,
		},
		{
			name: "False User not set",
			fields: fields{
				user: nil,
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
				user:              User.New().SetEmail("my@example.com"),
				simulateErrorFlag: true,
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
			r := &Repository{
				user:              tt.fields.user,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotIsExist, err := r.IsUserExist(tt.args.user)
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

func TestRepository_Auth(t *testing.T) {
	type fields struct {
		user              UserInterface.User
		simulateErrorFlag bool
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
			name: "True",
			fields: fields{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantIsValid: true,
			wantErr:     false,
		},
		{
			name: "False user mismatch",
			fields: fields{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			args: args{
				user: User.New().SetEmail("another@example.com").SetPassword("myPassword"),
			},
			wantIsValid: false,
			wantErr:     false,
		},
		{
			name: "False password mismatch",
			fields: fields{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("anotherPassword"),
			},
			wantIsValid: false,
			wantErr:     false,
		},
		{
			name: "False User not set",
			fields: fields{
				user: nil,
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
				user:              User.New().SetEmail("my@example.com").SetPassword("myPassword"),
				simulateErrorFlag: true,
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
			r := &Repository{
				user:              tt.fields.user,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotIsValid, err := r.Auth(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("Auth() gotIsValid = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}
