package cmd

import (
	"fmt"
	"os"

	"github.com/duckbrain/shiboleet/actions"
	"github.com/duckbrain/shiboleet/services"
	"github.com/gobuffalo/logger"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const AppName = "shiboleet"

var RootCmd = &cobra.Command{
	Use: AppName,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return viper.Unmarshal(&p.Config)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		app := actions.App(p)
		app.Options.Name = AppName
		app.Options.Addr = p.Config.ListenAddr
		app.Options.Env = p.Environment
		app.Options.SessionName = AppName
		return app.Serve()
	},
}

var p = services.Provider{
	Config: services.Config{
		Environment: "development",
		ListenAddr:  "0.0.0.0:3000",
		LogLevel:    logger.DebugLevel,
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
