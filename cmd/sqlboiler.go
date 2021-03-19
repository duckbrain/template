package cmd

import (
	"github.com/duckbrain/shiboleet/lib/runner"
	"github.com/spf13/cobra"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)

var BoilCmd = &cobra.Command{
	Use:  "boil",
	RunE: runner.CmdScript("sqlboiler"),
}

func init() {
	GenerateCmd.AddCommand(BoilCmd)

	afterMigrateHook = BoilCmd.RunE
}
