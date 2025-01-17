package cmd

import (
	out "github.com/LBJ-1/go-gitlab-client/cli/output"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getProjectPipelineCmd)
}

var getProjectPipelineCmd = &cobra.Command{
	Use:     resourceCmd("project-pipeline", "project-pipeline"),
	Aliases: []string{"pp"},
	Short:   "Get project pipeline details",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project-pipeline", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project's pipeline (project id: %s, pipeline id: %s)…", ids["project_id"], ids["pipeline_id"])

		loader.Start()
		pipeline, meta, err := client.ProjectPipeline(ids["project_id"], ids["pipeline_id"])
		loader.Stop()
		if err != nil {
			return err
		}

		out.Pipeline(output, outputFormat, pipeline)

		printMeta(meta, false)

		return nil
	},
}
