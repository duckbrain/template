package main

import (
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
		return api.Generate(cfg)
		// api.AddPlugin(gbgen.NewConvertPlugin(
		// 	output,   // directory where convert.go, convert_input.go and preload.go should live
		// 	backend,  // directory where sqlboiler files are put
		// 	frontend, // directory where gqlgen models live
		// 	gbgen.ConvertPluginConfig{
		// 		// UseReflectWorkaroundForSubModelFilteringInPostgresIssue25: true, // see issue #25 on GitHub
		// 	},
		// )),
		// api.AddPlugin(gbgen.NewResolverPlugin(
		// 	output,
		// 	backend,
		// 	frontend,
		// 	gbgen.ResolverPluginConfig{
		// 		// See example for AuthorizationScopes here: https://github.com/web-ridge/gqlgen-sqlboiler-examples/blob/main/social-network/convert_plugin.go#L66
		// 	},
		// )),

	},
}

func main() {
	runner.Main(GqlGenCmd)
}
