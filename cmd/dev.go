package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DevCmd = &cobra.Command{
	Use:     "dev",
	Aliases: []string{"watch"},
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(p.Config)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(DevCmd)
}
