package SessionIdGeneratorMock

import (
	"github.com/AntonParaskiv/my-life-assistant-back/domain/Session/Session"
	"github.com/AntonParaskiv/my-life-assistant-back/domain/User/User"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantG *Generator
	}{
		{
			name:  "Success",
			wantG: &Generator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotG := New(); !reflect.DeepEqual(gotG, tt.wantG) {
				t.Errorf("New() = %v, want %v", gotG, tt.wantG)
			}
		})
	}
}

func TestGenerator_md5(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		wantHash string
	}{
		{
			name: "Success",
			args: args{
				input: "my@example.com",
			},
			wantHash: "dfaa445d48f45b55bd695d89e593063c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{}
			if gotHash := g.md5(tt.args.input); gotHash != tt.wantHash {
				t.Errorf("md5() = %v, want %v", gotHash, tt.wantHash)
			}
		})
	}
}

func TestGenerator_GenerateIdFromString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		args   args
		wantId string
	}{
		{
			name: "Success",
			args: args{
				input: "my@example.com",
			},
			wantId: "dfaa445d48f45b55bd695d89e593063c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{}
			if gotId := g.GenerateIdFromString(tt.args.input); gotId != tt.wantId {
				t.Errorf("GenerateIdFromString() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestGenerator_Generate(t *testing.T) {
	type args struct {
		session *Session.Session
	}
	tests := []struct {
		name string
		args args
		want *Session.Session
	}{
		{
			name: "Success",
			args: args{
				session: Session.New().SetUser(User.New().SetEmail("my@example.com")),
			},
			want: Session.New().
				SetUser(User.New().SetEmail("my@example.com")).
				SetId("dfaa445d48f45b55bd695d89e593063c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{}
			if got := g.Generate(tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
