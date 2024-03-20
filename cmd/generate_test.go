package cmd_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/anibmurthy/htmgenie/cmd"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid Param",
			input:   "./testdata/dummy.md",
			wantErr: false,
		},
		{
			name:    "unsupported file",
			input:   "./testdata/dummy.txt",
			wantErr: true,
		},
		{
			name:    "Invalid path 2",
			input:   "./testdata//dummy.md",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rootCmd := cmd.RootCmd()
			buf := new(bytes.Buffer)
			rootCmd.SetOutput(buf)
			rootCmd.AddCommand(cmd.GenerateCommand())
			rootCmd.SetArgs([]string{
				"generate",
				"--file",
				test.input,
			})

			if err := rootCmd.Execute(); err != nil && !test.wantErr {
				t.Errorf("Exception occurred: %v", err)
			}
		})
	}
}

func TestGenerateCmdRun(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		got   string
	}{
		{
			name:  "Valid file",
			input: "./testdata/dummy.md",
			want:  "./testdata/expect1.html",
			got:   "./testdata/dummy.html",
		},
		{
			name:  "Valid file 2",
			input: "./testdata/dummy2.md",
			want:  "./testdata/expect2.html",
			got:   "./testdata/dummy2.html",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Setenv("HTMGENIE_OPATH", "./testdata")
			rootCmd := cmd.RootCmd()
			buf := new(bytes.Buffer)
			rootCmd.SetOutput(buf)
			rootCmd.AddCommand(cmd.GenerateCommand())
			rootCmd.SetArgs([]string{
				"generate",
				"--file",
				test.input,
			})

			if err := rootCmd.Execute(); err != nil {
				t.Errorf("Exception occurred: %v", err)
			}

			wb, err := os.ReadFile(test.want)
			if err != nil {
				t.Errorf("Exception occurred: %v", err)
			}

			gb, err := os.ReadFile(test.got)
			if err != nil {
				t.Errorf("Exception occurred: %v", err)
			}

			want := string(wb)
			got := string(gb)

			if compare(string(want), string(got)) {
				t.Errorf("Expected %v; got %v", test.want, got)
			}
		})
	}
}

func compare(s1 string, s2 string) bool {
	s1 = strings.ReplaceAll(s1, "\n", "")
	s2 = strings.ReplaceAll(s2, "\n", "")

	return s1 != s2
}
