package cmd

import (
	"testing"

	"github.com/LBJ-1/go-gitlab-client/test"
)

func TestGetCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		{
			[]string{"get"},
			nil,
			//configs["default"],
			"get_help",
			false,
			nil,
		},
		{
			[]string{"get", "--help"},
			nil,
			//configs["default"],
			"get_help",
			false,
			nil,
		},
	})
}
