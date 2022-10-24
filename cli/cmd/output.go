package cmd

import (
	out "github.com/LBJ-1/go-gitlab-client/cli/output"
	"github.com/LBJ-1/go-gitlab-client/gitlab"
)

func printMeta(meta *gitlab.ResponseMeta, withPagination bool) {
	if verbose {
		out.Meta(meta, withPagination)
	}
}
