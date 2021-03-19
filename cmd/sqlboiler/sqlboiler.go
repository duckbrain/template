package main

import (
	"github.com/duckbrain/shiboleet/lib/config"
	"github.com/duckbrain/shiboleet/lib/runner"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boilingcore"
	"github.com/volatiletech/sqlboiler/v4/drivers"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
	"github.com/volatiletech/sqlboiler/v4/importers"
)

const AppName = "shiboleet"

var conf = config.Default(AppName)

var BoilCmd = &cobra.Command{
	Use: "boil",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return config.Load(AppName, &conf)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		boilDriver := conf.Database.Driver
		if boilDriver == "" {
			boilDriver = conf.Database.Dialect
		}
		switch boilDriver {
		case "postgres":
			boilDriver = "psql"
		}

		config := &boilingcore.Config{
			DriverName: boilDriver,
			DriverConfig: drivers.Config{
				"dbname":    conf.Database.Database,
				"user":      conf.Database.User,
				"password":  conf.Database.Password,
				"host":      conf.Database.Host,
				"sslmode":   "disable",
				"blacklist": []string{"schema_migration"},
			},
			OutFolder: "models",
			PkgName:   "models",
			Imports:   importers.NewDefaultImports(),
			Wipe:      true,
		}
		core, err := boilingcore.New(config)
		if err != nil {
			return errors.Wrap(err, "new core")
		}
		return core.Run()
	},
}

func main() {
	runner.Main(BoilCmd)
}
