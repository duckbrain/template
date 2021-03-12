package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boilingcore"
	"github.com/volatiletech/sqlboiler/v4/drivers"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)

var BoilCmd = &cobra.Command{
	Use: "boil",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := &boilingcore.Config{
			DriverName: "pqsl",
			DriverConfig: drivers.Config{
				"dbname": "shiboleet",
			},
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
}
