package cmd

import (
	out "github.com/LBJ-1/go-gitlab-client/cli/output"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getUserCmd)
}

var getUserCmd = &cobra.Command{
	Use:     resourceCmd("user", "user"),
	Aliases: []string{"u"},
	Short:   "Get a single user",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "user", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching user (id: %s)…", ids["user_id"])

		loader.Start()
		user, meta, err := client.User(ids["user_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		out.User(output, outputFormat, user)

		printMeta(meta, false)

		relatedCommands([]*relatedCommand{
			newRelatedCommand(listUserSshKeysCmd, map[string]string{
				"user_id": ids["user_id"],
			}),
		})

		return nil
	},
}
