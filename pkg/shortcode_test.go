package pkg

import (
	"reflect"
	"testing"
)

func TestParseShortCode(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{{
		name: "org/repo format",
		args: args{[]string{"org/repo"}},
		want: []string{"https://github.com/org/repo", "repo"},
	}, {
		name: "'org/repo output' format",
		args: args{[]string{"org/repo", "dir"}},
		want: []string{"https://github.com/org/repo", "dir"},
	}, {
		name: "unexpected format",
		args: args{[]string{}},
		want: []string{},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseShortCode(tt.args.args, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseShortCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
