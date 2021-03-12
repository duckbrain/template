package cmd

import (
	"github.com/gobuffalo/buffalo/buffalo/cmd/generate"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate application components",
	Aliases: []string{"g"},
}

func init() {
	GenerateCmd.AddCommand(generate.ResourceCmd)
	GenerateCmd.AddCommand(generate.ActionCmd)
	GenerateCmd.AddCommand(generate.TaskCmd)
	GenerateCmd.AddCommand(generate.MailCmd)

	RootCmd.AddCommand(GenerateCmd)
}
