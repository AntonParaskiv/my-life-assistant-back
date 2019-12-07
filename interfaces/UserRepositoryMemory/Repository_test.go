package UserRepositoryMemory

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserList"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserListInterface"
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

func TestRepository_SetUserList(t *testing.T) {
	type fields struct {
		userList UserListInterface.UserListInterface
	}
	type args struct {
		userList UserListInterface.UserListInterface
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
				userList: UserList.New().Add(User.New().SetEmail("my@example.com")),
			},
			want: &Repository{
				userList: UserList.New().Add(User.New().SetEmail("my@example.com")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				userList: tt.fields.userList,
			}
			if got := r.SetUserList(tt.args.userList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_IsUserExist(t *testing.T) {
	type fields struct {
		userList UserListInterface.UserListInterface
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
			name: "Exist",
			fields: fields{
				userList: UserList.New().
					Add(User.New().SetEmail("first@user.com")).
					Add(User.New().SetEmail("second@user.com")).
					Add(User.New().SetEmail("third@user.com")),
			},
			args: args{
				user: User.New().SetEmail("second@user.com"),
			},
			wantIsExist: true,
			wantErr:     false,
		},
		{
			name: "NotExist",
			fields: fields{
				userList: UserList.New().
					Add(User.New().SetEmail("first@user.com")).
					Add(User.New().SetEmail("second@user.com")).
					Add(User.New().SetEmail("third@user.com")),
			},
			args: args{
				user: User.New().SetEmail("fourth@user.com"),
			},
			wantIsExist: false,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				userList: tt.fields.userList,
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

func TestRepository_addUser(t *testing.T) {
	type fields struct {
		userList *UserList.List
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
			name: "Success",
			fields: fields{
				userList: UserList.New().
					Add(User.New().SetEmail("first@user.com")),
			},
			args: args{
				user: User.New().SetEmail("second@user.com"),
			},
			want: &Repository{
				userList: UserList.New().
					Add(User.New().SetEmail("first@user.com")).
					Add(User.New().SetEmail("second@user.com")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				userList: tt.fields.userList,
			}
			if got := r.addUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_AddUser(t *testing.T) {
	type fields struct {
		userList *UserList.List
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
			name: "Error User Exist",
			fields: fields{
				userList: UserList.New().
					Add(User.New().SetEmail("first@user.com")),
			},
			args: args{
				user: User.New().SetEmail("first@user.com"),
			},
			wantErr: true,
			wantRepository: &Repository{
				userList: UserList.New().
					Add(User.New().SetEmail("first@user.com")),
			},
		},
		{
			name: "Success ",
			fields: fields{
				userList: UserList.New().
					Add(User.New().SetEmail("first@user.com")),
			},
			args: args{
				user: User.New().SetEmail("second@user.com"),
			},
			wantErr: false,
			wantRepository: &Repository{
				userList: UserList.New().
					Add(User.New().SetEmail("first@user.com")).
					Add(User.New().SetEmail("second@user.com")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				userList: tt.fields.userList,
			}
			if err := r.AddUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("Repository = %v, want %v", r, tt.wantRepository)
			}
		})
	}
}

func TestRepository_Auth(t *testing.T) {
	type fields struct {
		userList UserListInterface.UserListInterface
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
				userList: UserList.New().
					Add(User.New().SetEmail("my@example.com").SetPassword("myPassword")),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantIsValid: true,
			wantErr:     false,
		},
		{
			name: "User Not Exist",
			fields: fields{
				userList: UserList.New(),
			},
			args: args{
				user: User.New().SetEmail("my@example.com").SetPassword("myPassword"),
			},
			wantIsValid: false,
			wantErr:     true,
		},
		{
			name: "Password Not Match",
			fields: fields{
				userList: UserList.New().
					Add(User.New().SetEmail("my@example.com").SetPassword("anotherPassword")),
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
				userList: tt.fields.userList,
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
