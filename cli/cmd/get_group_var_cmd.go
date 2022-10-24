package cmd

import (
	out "github.com/LBJ-1/go-gitlab-client/cli/output"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getGroupVarCmd)
}

var getGroupVarCmd = &cobra.Command{
	Use:     resourceCmd("group-var", "group-var"),
	Aliases: []string{"gv"},
	Short:   "Get the details of a group's specific variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group-var", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching group variable (group id: %s, key: %s)…", ids["group_id"], ids["var_key"])

		loader.Start()
		variable, meta, err := client.GroupVariable(ids["group_id"], ids["var_key"])
		loader.Stop()
		if err != nil {
			return err
		}

		out.Variable(output, outputFormat, variable)

		printMeta(meta, false)

		return nil
	},
}
