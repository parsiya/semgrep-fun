package fun

import (
	"testing"
)

func TestTopLevelPackageOld(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{"github.com/foo/bar"},
			want: "bar",
		},
		{
			name: "no slash",
			args: args{"github.com"},
			want: "github.com",
		},
		{
			name: "empty",
			args: args{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopLevelPackageOld(tt.args.input); got != tt.want {
				t.Errorf("TopLevelPackageOld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopLevelPackage(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{"github.com/foo/bar"},
			want: "bar",
		},
		{
			name: "no slash",
			args: args{"github.com"},
			want: "github.com",
		},
		{
			name: "empty",
			args: args{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopLevelPackage(tt.args.input); got != tt.want {
				t.Errorf("TopLevelPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}
