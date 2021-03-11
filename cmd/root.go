package cmd

import (
	"fmt"
	"os"

	"github.com/duckbrain/shiboleet/actions"
	"github.com/duckbrain/shiboleet/services"
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/pop/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const AppName = "shiboleet"

var RootCmd = &cobra.Command{
	Use: AppName,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := viper.Unmarshal(&p.Config)
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
		Environment: "development",
		ListenAddr:  "0.0.0.0:3000",
		LogLevel:    logger.DebugLevel,
		Database: &pop.ConnectionDetails{
			Dialect:  "postgres",
			Host:     "localhost",
			User:     "postgres",
			Database: AppName,
		},
		Assets: struct {
			Type                  string
			ManifestPath          string
			DevelopmentServerPath string
		}{
			Type: "vite",
		},
	},
}

func init() {
	RootCmd.Flags().String("addr", "", "Listen address for the application")
	must(viper.BindPFlag("ListenAddr", RootCmd.Flag("addr")))

	viper.SetConfigName(AppName)
	viper.SetEnvPrefix(AppName)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/" + AppName)
	viper.AutomaticEnv()
	defaults := map[string]interface{}{}
	must(mapstructure.Decode(p.Config, &defaults))
	for key, val := range defaults {
		viper.SetDefault(key, val)
	}
	viper.SetTypeByDefaultValue(true)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
