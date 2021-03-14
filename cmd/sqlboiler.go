package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boilingcore"
	"github.com/volatiletech/sqlboiler/v4/drivers"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
	"github.com/volatiletech/sqlboiler/v4/importers"
)

var BoilCmd = &cobra.Command{
	Use: "boil",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := &boilingcore.Config{
			DriverName: "psql",
			DriverConfig: drivers.Config{
				"dbname":    "shiboleet",
				"user":      "postgres",
				"password":  "postgres",
				"host":      "localhost",
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

func init() {
	GenerateCmd.AddCommand(BoilCmd)

	afterMigrateHook = BoilCmd.RunE
}
