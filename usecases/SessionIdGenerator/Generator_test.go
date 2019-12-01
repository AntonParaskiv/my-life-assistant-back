package SessionIdGenerator

import (
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
