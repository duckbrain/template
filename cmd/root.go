package cmd

import (
	"github.com/duckbrain/shiboleet/actions"
	"github.com/duckbrain/shiboleet/lib/config"
	"github.com/duckbrain/shiboleet/services"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const AppName = "shiboleet"

var RootCmd = &cobra.Command{
	Use: AppName,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := config.Load(AppName, &p.Config)
		if err != nil {
			return errors.Wrap(err, "config")
		}
		return p.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return actions.NewApp(p).Serve()
	},
}

var MigrationBox packd.Box = packr.New("migrations", "../migrations")
var p = services.Provider{
	AppName: AppName,
	Config: services.Config{
		Config: config.Default(AppName),
	},
}.Must()

func init() {
	RootCmd.Flags().String("addr", "", "Listen address for the application")
	must(viper.BindPFlag("ListenAddr", RootCmd.Flag("addr")))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
