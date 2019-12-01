package Session

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantS *Session
	}{
		{
			name:  "Success",
			wantS: &Session{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := New(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("New() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestSession_SetUser(t *testing.T) {
	type fields struct {
		user *User.User
	}
	type args struct {
		user *User.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Session
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				user: User.New().SetEmail("email@example.com"),
			},
			want: &Session{
				user: User.New().SetEmail("email@example.com"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				user: tt.fields.user,
			}
			if got := s.SetUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_SetId(t *testing.T) {
	type fields struct {
		id string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Session
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				id: "myId",
			},
			want: &Session{
				id: "myId",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				id: tt.fields.id,
			}
			if got := s.SetId(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_User(t *testing.T) {
	type fields struct {
		user *User.User
	}
	tests := []struct {
		name     string
		fields   fields
		wantUser *User.User
	}{
		{
			name: "Success",
			fields: fields{
				user: User.New().SetEmail("email@example.com"),
			},
			wantUser: User.New().SetEmail("email@example.com"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				user: tt.fields.user,
			}
			if gotUser := s.User(); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("User() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestSession_Id(t *testing.T) {
	type fields struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		wantId string
	}{
		{
			name: "Success",
			fields: fields{
				id: "myId",
			},
			wantId: "myId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				id: tt.fields.id,
			}
			if gotId := s.Id(); gotId != tt.wantId {
				t.Errorf("Id() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
