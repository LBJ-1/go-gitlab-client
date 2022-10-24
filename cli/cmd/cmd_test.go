package cmd

import (
	"os"
	"testing"

	"github.com/LBJ-1/go-gitlab-client/test"
)

func TestMain(m *testing.M) {
	test.BuildCli()

	os.Exit(m.Run())
}
