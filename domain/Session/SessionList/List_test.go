package SessionList

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
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
			if gotL := New(); !reflect.DeepEqual(gotL, tt.wantList) {
				t.Errorf("New() = %v, want %v", gotL, tt.wantList)
			}
		})
	}
}

func TestList_Add(t *testing.T) {
	type args struct {
		session *Session.Session
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
				Session.New().SetUser(User.New().SetEmail("first@user.com")),
			},
			args: args{
				Session.New().SetUser(User.New().SetEmail("second@user.com")),
			},
			want: &List{
				Session.New().SetUser(User.New().SetEmail("first@user.com")),
				Session.New().SetUser(User.New().SetEmail("second@user.com")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Add(tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_AddSession(t *testing.T) {
	type args struct {
		session *Session.Session
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
				Session.New().SetUser(User.New().SetEmail("first@user.com")),
			},
			args: args{
				Session.New().SetUser(User.New().SetEmail("second@user.com")),
			},
			wantErr: false,
			wantList: List{
				Session.New().SetUser(User.New().SetEmail("first@user.com")),
				Session.New().SetUser(User.New().SetEmail("second@user.com")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.list.AddSession(tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("AddSession() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.list, tt.wantList) {
				t.Errorf("list = %v, want %v", tt.list, tt.wantList)
			}
		})
	}
}

func TestList_searchSessionById(t *testing.T) {
	type args struct {
		id string
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
				Session.New().SetId("firstId"),
				Session.New().SetId("secondId"),
				Session.New().SetId("thirdId"),
			},
			args: args{
				id: "fourthId",
			},
			wantKey: -1,
		},
		{
			name: "Exist",
			list: List{
				Session.New().SetId("firstId"),
				Session.New().SetId("secondId"),
				Session.New().SetId("thirdId"),
			},
			args: args{
				id: "secondId",
			},
			wantKey: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKey := tt.list.searchSessionById(tt.args.id); gotKey != tt.wantKey {
				t.Errorf("searchSessionById() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

func TestList_Len(t *testing.T) {
	tests := []struct {
		name          string
		list          List
		wantListength int
	}{
		{
			name:          "Zero",
			list:          List{},
			wantListength: 0,
		},
		{
			name: "Three",
			list: List{
				Session.New(),
				Session.New(),
				Session.New(),
			},
			wantListength: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLength := tt.list.Len(); gotLength != tt.wantListength {
				t.Errorf("Len() = %v, want %v", gotLength, tt.wantListength)
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
				Session.New(),
			},
			wantKey: 0,
		},
		{
			name: "Three",
			list: List{
				Session.New(),
				Session.New(),
				Session.New(),
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
				Session.New(),
				Session.New(),
				Session.New(),
			},
			args: args{
				key: 1,
			},
			wantIsValid: true,
		},
		{
			name: "InValid < 0",
			list: List{
				Session.New(),
				Session.New(),
				Session.New(),
			},
			args: args{
				key: -1,
			},
			wantIsValid: false,
		},
		{
			name: "InValid > lastKey",
			list: List{
				Session.New(),
				Session.New(),
				Session.New(),
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

func TestList_getSessionByKey(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name        string
		list        List
		args        args
		wantSession *Session.Session
	}{
		{
			name: "Exist",
			list: List{
				Session.New().SetUser(User.New().SetEmail("first@user.com")),
				Session.New().SetUser(User.New().SetEmail("second@user.com")),
				Session.New().SetUser(User.New().SetEmail("third@user.com")),
			},
			args: args{
				key: 1,
			},
			wantSession: Session.New().SetUser(User.New().SetEmail("second@user.com")),
		},
		{
			name: "NotExist < 0",
			list: List{
				Session.New().SetUser(User.New().SetEmail("first@user.com")),
				Session.New().SetUser(User.New().SetEmail("second@user.com")),
				Session.New().SetUser(User.New().SetEmail("third@user.com")),
			},
			args: args{
				key: -1,
			},
			wantSession: nil,
		},
		{
			name: "NotExist > lastKey",
			list: List{
				Session.New().SetUser(User.New().SetEmail("first@user.com")),
				Session.New().SetUser(User.New().SetEmail("second@user.com")),
				Session.New().SetUser(User.New().SetEmail("third@user.com")),
			},
			args: args{
				key: 3,
			},
			wantSession: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSession := tt.list.getSessionByKey(tt.args.key); !reflect.DeepEqual(gotSession, tt.wantSession) {
				t.Errorf("getSessionByKey() = %v, want %v", gotSession, tt.wantSession)
			}
		})
	}
}

func TestList_GetSessionById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name        string
		list        List
		args        args
		wantSession *Session.Session
	}{
		{
			name: "Success",
			list: List{
				Session.New().SetId("firstId"),
				Session.New().SetId("secondId"),
				Session.New().SetId("thirdId"),
			},
			args: args{
				id: "secondId",
			},
			wantSession: Session.New().SetId("secondId"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSession := tt.list.GetSessionById(tt.args.id); !reflect.DeepEqual(gotSession, tt.wantSession) {
				t.Errorf("GetSessionById() = %v, want %v", gotSession, tt.wantSession)
			}
		})
	}
}

func TestList_IsSessionIdExist(t *testing.T) {
	type args struct {
		session *Session.Session
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
				Session.New().SetId("firstId"),
				Session.New().SetId("secondId"),
				Session.New().SetId("thirdId"),
			},
			args: args{
				session: Session.New().SetId("secondId"),
			},
			wantIsExist: true,
		},
		{
			name: "NotExist",
			list: List{
				Session.New().SetId("firstId"),
				Session.New().SetId("secondId"),
				Session.New().SetId("thirdId"),
			},
			args: args{
				session: Session.New().SetId("fourthId"),
			},
			wantIsExist: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsExist := tt.list.IsSessionIdExist(tt.args.session); gotIsExist != tt.wantIsExist {
				t.Errorf("IsSessionIdExist() = %v, want %v", gotIsExist, tt.wantIsExist)
			}
		})
	}
}
