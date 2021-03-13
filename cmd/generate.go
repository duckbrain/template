package cmd

import (
	buffalogen "github.com/gobuffalo/buffalo/buffalo/cmd/generate"
	sodagen "github.com/gobuffalo/pop/v5/soda/cmd/generate"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate application components",
	Aliases: []string{"g"},
}

func init() {
	GenerateCmd.AddCommand(buffalogen.ResourceCmd)
	GenerateCmd.AddCommand(buffalogen.ActionCmd)
	GenerateCmd.AddCommand(buffalogen.TaskCmd)
	GenerateCmd.AddCommand(buffalogen.MailCmd)

	GenerateCmd.AddCommand(sodagen.ConfigCmd)
	GenerateCmd.AddCommand(sodagen.FizzCmd)
	GenerateCmd.AddCommand(sodagen.SQLCmd)
	GenerateCmd.AddCommand(sodagen.ModelCmd)

	RootCmd.AddCommand(GenerateCmd)
}
