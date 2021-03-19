package config

import (
	"github.com/duckbrain/shiboleet/lib/assets"
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/pop/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func Default(AppName string) Config {
	return Config{
		Environment: "development",
		ListenAddr:  "0.0.0.0:3000",
		LogLevel:    logger.DebugLevel,
		Database: &pop.ConnectionDetails{
			Dialect:  "postgres",
			Host:     "localhost",
			User:     "postgres",
			Database: AppName,
		},
		Assets: assets.Config{
			Type: "vite",
		},
	}
}

type Config struct {
	// Environment the application will be running in, eg: "development" and "production"
	Environment string

	// Server options
	ListenAddr string // Address the web server will listen on
	Host       string // Absolute URL the application will be accessed through

	LogLevel logger.Level

	Database *pop.ConnectionDetails

	Assets assets.Config
}

func Load(AppName string, config interface{}) error {
	viper.SetConfigName(AppName)
	viper.SetEnvPrefix(AppName)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/" + AppName)
	viper.AutomaticEnv()
	defaults := map[string]interface{}{}
	err := mapstructure.Decode(config, &defaults)
	if err != nil {
		return errors.Wrap(err, "mapstructure")
	}
	for key, val := range defaults {
		viper.SetDefault(key, val)
	}
	viper.SetTypeByDefaultValue(true)
	err = viper.Unmarshal(config)
	if err != nil {
		return errors.Wrap(err, "config")
	}

	if i, ok := config.(interface{ Init() error }); ok {
		err = i.Init()
	}

	return err
}
