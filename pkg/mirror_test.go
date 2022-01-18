package pkg

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestUseMirror(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		args   args
		verify func(*testing.T, []string)
	}{{
		name: "normal case",
		args: args{
			[]string{"clone", "https://github.com/a/b"},
		},
		verify: func(t *testing.T, args []string) {
			assert.Equal(t, args[0], "clone")
			assert.Equal(t, args[1], "https://github.com.cnpmjs.org/a/b")
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMirror(tt.args.args)
			tt.verify(t, tt.args.args)
		})
	}
}
