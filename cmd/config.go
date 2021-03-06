package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:  "config",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(p.Config)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(ConfigCmd)
}
