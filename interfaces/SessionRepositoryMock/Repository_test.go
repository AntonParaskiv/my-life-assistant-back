package SessionRepositoryMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/SessionInterface"
	"github.com/AntonParaskiv/my-life-assistant-back/interfaces/SessionRepositoryInterface"
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

func TestRepository_SetSession(t *testing.T) {
	type fields struct {
		session SessionInterface.Session
	}
	type args struct {
		session SessionInterface.Session
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SessionRepositoryInterface.Repository
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				session: Session.New(),
			},
			want: &Repository{
				session: Session.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				session: tt.fields.session,
			}
			if got := r.SetSession(tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Session(t *testing.T) {
	type fields struct {
		session           SessionInterface.Session
		simulateErrorFlag bool
	}
	tests := []struct {
		name        string
		fields      fields
		wantSession SessionInterface.Session
	}{
		{
			name: "Success",
			fields: fields{
				session: Session.New(),
			},
			wantSession: Session.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				session:           tt.fields.session,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if gotSession := r.Session(); !reflect.DeepEqual(gotSession, tt.wantSession) {
				t.Errorf("Session() = %v, want %v", gotSession, tt.wantSession)
			}
		})
	}
}

func TestRepository_AddSession(t *testing.T) {
	type fields struct {
		session           SessionInterface.Session
		simulateErrorFlag bool
	}
	type args struct {
		session SessionInterface.Session
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
				session: Session.New(),
			},
			wantErr: false,
			wantRepository: &Repository{
				session: Session.New(),
			},
		},
		{
			name: "Error",
			fields: fields{
				simulateErrorFlag: true,
			},
			args: args{
				session: Session.New(),
			},
			wantErr: true,
			wantRepository: &Repository{
				session: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				session:           tt.fields.session,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := r.AddSession(tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("AddSession() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("Repository= %v, want %v", r, tt.wantRepository)
			}
		})
	}
}

func TestRepository_IsSessionIdExist(t *testing.T) {
	type fields struct {
		session           SessionInterface.Session
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
			name: "True",
			fields: fields{
				session: Session.New().SetId("myId"),
			},
			args: args{
				session: Session.New().SetId("myId"),
			},
			wantIsExist: true,
			wantErr:     false,
		},
		{
			name: "False",
			fields: fields{
				session: Session.New().SetId("myId"),
			},
			args: args{
				session: Session.New().SetId("anotherId"),
			},
			wantIsExist: false,
			wantErr:     false,
		},
		{
			name: "False Session nil",
			fields: fields{
				session: nil,
			},
			args: args{
				session: Session.New().SetId("anotherId"),
			},
			wantIsExist: false,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				session:           tt.fields.session,
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
