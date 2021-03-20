package main

import (
	"github.com/duckbrain/buffalo-gqlgen/gqlgen/plugin"
	"github.com/duckbrain/shiboleet/lib/runner"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

var GqlGenCmd = &cobra.Command{
	Use: "gqlgen",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfigFromDefaultLocations()
		if err != nil {
			return errors.Wrap(err, "load config")
		}
		return api.Generate(cfg, api.AddPlugin(plugin.Scalar{}))

	},
}

func main() {
	runner.Main(GqlGenCmd)
}
