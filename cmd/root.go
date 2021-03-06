package cmd

import (
	"fmt"
	"os"

	"github.com/duckbrain/shiboleet/actions"
	"github.com/duckbrain/shiboleet/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use: "shiboleet",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return viper.Unmarshal(p)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return actions.App(p).Serve()
	},
}

var p = services.Provider{}

func Run() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
