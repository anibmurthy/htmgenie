package cmd_test

import (
	"testing"

	"github.com/anibmurthy/htmgenie/cmd"
)

func Test_ExecuteCommand(t *testing.T) {
	root := cmd.RootCmd()

	if err := cmd.Execute(root); err != nil {
		t.Error("Root command failed", err)
	}
}
