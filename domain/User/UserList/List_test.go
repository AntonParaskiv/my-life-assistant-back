package UserList

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		wantList *List
	}{
		{
			name:     "Success",
			wantList: &List{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotList := New(); !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("New() = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}

func TestList_Add(t *testing.T) {
	type args struct {
		user *User.User
	}
	tests := []struct {
		name string
		list List
		args args
		want *List
	}{
		{
			name: "Success",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
			},
			args: args{
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
			},
			want: &List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Add(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_AddUser(t *testing.T) {
	type args struct {
		user *User.User
	}
	tests := []struct {
		name     string
		list     List
		args     args
		wantErr  bool
		wantList List
	}{
		{
			name: "Success",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
			},
			args: args{
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
			},
			wantList: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.list.AddUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.list, tt.wantList) {
				t.Errorf("list = %v, want %v", tt.list, tt.wantList)
			}
		})
	}
}

func TestList_searchUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		list    List
		args    args
		wantKey int
	}{
		{
			name: "Not Exist",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
				User.New().
					SetEmail("third@user.com").
					SetPassword("thirdPassword"),
			},
			args: args{
				email: "fourth@user.com",
			},
			wantKey: -1,
		},
		{
			name: "Exist",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
				User.New().
					SetEmail("third@user.com").
					SetPassword("thirdPassword"),
			},
			args: args{
				email: "second@user.com",
			},
			wantKey: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKey := tt.list.searchUserByEmail(tt.args.email); gotKey != tt.wantKey {
				t.Errorf("searchUserByEmail() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

func TestList_Len(t *testing.T) {
	tests := []struct {
		name       string
		list       List
		wantLength int
	}{
		{
			name:       "Zero",
			list:       List{},
			wantLength: 0,
		},
		{
			name: "Three",
			list: List{
				User.New(),
				User.New(),
				User.New(),
			},
			wantLength: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLength := tt.list.Len(); gotLength != tt.wantLength {
				t.Errorf("Len() = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}

func TestList_lastKey(t *testing.T) {
	tests := []struct {
		name    string
		list    List
		wantKey int
	}{
		{
			name:    "Zero",
			list:    List{},
			wantKey: -1,
		},

		{
			name: "One",
			list: List{
				User.New(),
			},
			wantKey: 0,
		},
		{
			name: "Three",
			list: List{
				User.New(),
				User.New(),
				User.New(),
			},
			wantKey: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKey := tt.list.lastKey(); gotKey != tt.wantKey {
				t.Errorf("lastKey() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

func TestList_isKeyValid(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name        string
		list        List
		args        args
		wantIsValid bool
	}{
		{
			name: "Valid",
			list: List{
				User.New(),
				User.New(),
				User.New(),
			},
			args: args{
				key: 1,
			},
			wantIsValid: true,
		},
		{
			name: "InValid < 0",
			list: List{
				User.New(),
				User.New(),
				User.New(),
			},
			args: args{
				key: -1,
			},
			wantIsValid: false,
		},
		{
			name: "InValid > lastKey",
			list: List{
				User.New(),
				User.New(),
				User.New(),
			},
			args: args{
				key: 3,
			},
			wantIsValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsValid := tt.list.isKeyValid(tt.args.key); gotIsValid != tt.wantIsValid {
				t.Errorf("isKeyValid() = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}

func TestList_getUserByKey(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name     string
		list     List
		args     args
		wantUser *User.User
	}{
		{
			name: "Exist",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
				User.New().
					SetEmail("third@user.com").
					SetPassword("thirdPassword"),
			},
			args: args{
				key: 1,
			},
			wantUser: User.New().
				SetEmail("second@user.com").
				SetPassword("secondPassword"),
		},
		{
			name: "NotExist < 0",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
				User.New().
					SetEmail("third@user.com").
					SetPassword("thirdPassword"),
			},
			args: args{
				key: -1,
			},
			wantUser: nil,
		},
		{
			name: "NotExist > lastKey",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
				User.New().
					SetEmail("third@user.com").
					SetPassword("thirdPassword"),
			},
			args: args{
				key: 3,
			},
			wantUser: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUser := tt.list.getUserByKey(tt.args.key); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("getUserByKey() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestList_GetUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name     string
		list     List
		args     args
		wantUser *User.User
	}{
		{
			name: "Success",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
				User.New().
					SetEmail("third@user.com").
					SetPassword("thirdPassword"),
			},
			args: args{
				email: "second@user.com",
			},
			wantUser: User.New().
				SetEmail("second@user.com").
				SetPassword("secondPassword"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUser := tt.list.GetUserByEmail(tt.args.email); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUserByEmail() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestList_IsUserExist(t *testing.T) {
	type args struct {
		user *User.User
	}
	tests := []struct {
		name        string
		list        List
		args        args
		wantIsExist bool
	}{
		{
			name: "Exist",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
				User.New().
					SetEmail("third@user.com").
					SetPassword("thirdPassword"),
			},
			args: args{
				user: User.New().SetEmail("second@user.com"),
			},
			wantIsExist: true,
		},
		{
			name: "NotExist",
			list: List{
				User.New().
					SetEmail("first@user.com").
					SetPassword("firstPassword"),
				User.New().
					SetEmail("second@user.com").
					SetPassword("secondPassword"),
				User.New().
					SetEmail("third@user.com").
					SetPassword("thirdPassword"),
			},
			args: args{
				user: User.New().SetEmail("fourth@user.com"),
			},
			wantIsExist: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsExist := tt.list.IsUserExist(tt.args.user); gotIsExist != tt.wantIsExist {
				t.Errorf("IsUserExist() = %v, want %v", gotIsExist, tt.wantIsExist)
			}
		})
	}
}
