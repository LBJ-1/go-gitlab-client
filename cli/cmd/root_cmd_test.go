package cmd

import (
	"testing"

	"github.com/LBJ-1/go-gitlab-client/test"
)

func TestRootCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		{
			[]string{},
			nil,
			//configs["default"],
			"help",
			false,
			nil,
		},
		{
			[]string{"help"},
			nil,
			//configs["default"],
			"help",
			false,
			nil,
		},
	})
}
