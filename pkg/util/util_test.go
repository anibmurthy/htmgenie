package util_test

import (
	"os"
	"testing"

	"github.com/anibmurthy/htmgenie/pkg/util"
)

func TestGetOutPath(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		isEnv bool
	}{
		{
			name:  "simple path",
			input: "c:/input/path/that/isvalid.md",
			want:  "./isvalid.html",
		},
		{
			name:  "simple path - With Env",
			input: "c:/input/path/that/isvalid.md",
			want:  "c:/temp/htmlout/isvalid.html",
			isEnv: true,
		},
		{
			name:  "invalid path (without .md)",
			input: "c:/input/path/that/invalid",
			want:  "./invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isEnv {
				os.Setenv("HTMGENIE_OPATH", "c:/temp/htmlout")
			} else {
				os.Setenv("HTMGENIE_OPATH", "")
			}
			got := util.GetOutPath(test.input)

			if test.want != got {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}
