package SessionRepositoryMemory

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionList"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionListInterface"
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

func TestRepository_SetSessionList(t *testing.T) {
	type fields struct {
		sessionList SessionListInterface.SessionList
	}
	type args struct {
		sessionList SessionListInterface.SessionList
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
				sessionList: SessionList.New().
					Add(Session.New().SetId("myId")),
			},
			want: &Repository{
				sessionList: SessionList.New().
					Add(Session.New().SetId("myId")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				sessionList: tt.fields.sessionList,
			}
			if got := r.SetSessionList(tt.args.sessionList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSessionList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_IsSessionIdExist(t *testing.T) {
	type fields struct {
		sessionList       SessionListInterface.SessionList
		simulateErrorFlag bool
	}
	type args struct {
		session SessionInterface.Session
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
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")).
					Add(Session.New().SetId("secondId")).
					Add(Session.New().SetId("thirdId")),
			},
			args: args{
				session: Session.New().SetId("secondId"),
			},
			wantIsExist: true,
			wantErr:     false,
		},
		{
			name: "NotExist",
			fields: fields{
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")).
					Add(Session.New().SetId("secondId")).
					Add(Session.New().SetId("thirdId")),
			},
			args: args{
				session: Session.New().SetId("fourthId"),
			},
			wantIsExist: false,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				sessionList:       tt.fields.sessionList,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotIsExist, err := r.IsSessionIdExist(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSessionIdExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsExist != tt.wantIsExist {
				t.Errorf("IsSessionIdExist() gotIsExist = %v, want %v", gotIsExist, tt.wantIsExist)
			}
		})
	}
}

func TestRepository_addSession(t *testing.T) {
	type fields struct {
		sessionList *SessionList.List
	}
	type args struct {
		session *Session.Session
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
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")),
			},
			args: args{
				session: Session.New().SetId("secondId"),
			},
			want: &Repository{
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")).
					Add(Session.New().SetId("secondId")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				sessionList: tt.fields.sessionList,
			}
			if got := r.addSession(tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_AddSession(t *testing.T) {
	type fields struct {
		sessionList       *SessionList.List
		simulateErrorFlag bool
	}
	type args struct {
		session *Session.Session
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantRepository *Repository
	}{
		{
			name: "Error Simulated",
			fields: fields{
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")),
				simulateErrorFlag: true,
			},
			args: args{
				session: Session.New().SetId("firstId"),
			},
			wantErr: true,
			wantRepository: &Repository{
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")),
			},
		},
		{
			name: "Error Session Exist",
			fields: fields{
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")),
			},
			args: args{
				session: Session.New().SetId("firstId"),
			},
			wantErr: true,
			wantRepository: &Repository{
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")),
			},
		},
		{
			name: "Success ",
			fields: fields{
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")),
			},
			args: args{
				session: Session.New().SetId("secondId"),
			},
			wantErr: false,
			wantRepository: &Repository{
				sessionList: SessionList.New().
					Add(Session.New().SetId("firstId")).
					Add(Session.New().SetId("secondId")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				sessionList:       tt.fields.sessionList,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := r.AddSession(tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("AddSession() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("Repository = %v, want %v", r, tt.wantRepository)
			}
		})
	}
}
