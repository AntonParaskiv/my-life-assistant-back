package User

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/UserInterface"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantU *User
	}{
		{
			name:  "Success",
			wantU: &User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotU := New(); !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("New() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func TestUser_Email(t *testing.T) {
	type fields struct {
		email string
	}
	tests := []struct {
		name      string
		fields    fields
		wantEmail string
	}{
		{
			name: "Success",
			fields: fields{
				email: "my@example.com",
			},
			wantEmail: "my@example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				email: tt.fields.email,
			}
			if gotEmail := u.Email(); gotEmail != tt.wantEmail {
				t.Errorf("Email() = %v, want %v", gotEmail, tt.wantEmail)
			}
		})
	}
}

func TestUser_Password(t *testing.T) {
	type fields struct {
		password string
	}
	tests := []struct {
		name         string
		fields       fields
		wantPassword string
	}{
		{
			name: "Success",
			fields: fields{
				password: "myPassword",
			},
			wantPassword: "myPassword",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				password: tt.fields.password,
			}
			if gotPassword := u.Password(); gotPassword != tt.wantPassword {
				t.Errorf("Password() = %v, want %v", gotPassword, tt.wantPassword)
			}
		})
	}
}

func TestUser_SetEmail(t *testing.T) {
	type fields struct {
		email string
	}
	type args struct {
		email string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   UserInterface.User
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				email: "my@example.com",
			},
			want: &User{
				email: "my@example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				email: tt.fields.email,
			}
			if got := u.SetEmail(tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetPassword(t *testing.T) {
	type fields struct {
		password string
	}
	type args struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   UserInterface.User
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				password: "myPassword",
			},
			want: &User{
				password: "myPassword",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				password: tt.fields.password,
			}
			if got := u.SetPassword(tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
