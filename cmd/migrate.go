package cmd

import (
	"os"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/cobra"
)

// In debug, is overridden with sqlboiler's RunE.
var afterMigrateHook = func(cmd *cobra.Command, args []string) error {
	return nil
}

func MigrateCmd(c *pop.Connection, box packd.Box) *cobra.Command {
	m := &cobra.Command{
		Use: "migrate",
	}

	var mb pop.MigrationBox

	prepMB := func(cmd *cobra.Command, args []string) (err error) {
		mb, err = pop.NewMigrationBox(box, c)
		return
	}

	prepDB := func(cmd *cobra.Command, args []string) error {
		_ = pop.CreateDB(c)
		return prepMB(cmd, args)
	}

	postMigrate := func(cmd *cobra.Command, args []string) error {
		return afterMigrateHook(cmd, args)
	}

	var steps int

	withSteps := func(c *cobra.Command) *cobra.Command {
		c.Flags().IntVarP(&steps, "steps", "s", -1, "steps to move")
		return c
	}

	m.AddCommand(&cobra.Command{
		Use:     "create",
		PreRunE: prepMB,
		RunE: func(cmd *cobra.Command, args []string) error {
			return pop.CreateDB(c)
		},
	})
	m.AddCommand(&cobra.Command{
		Use:     "drop",
		PreRunE: prepMB,
		RunE: func(cmd *cobra.Command, args []string) error {
			return pop.DropDB(c)
		},
	})
	m.AddCommand(withSteps(&cobra.Command{
		Use:      "up",
		PreRunE:  prepDB,
		PostRunE: postMigrate,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := mb.UpTo(steps)
			// TODO log that not all the steps were done
			// if steps > 1 && s != steps {
			return err
		},
	}))
	m.AddCommand(withSteps(&cobra.Command{
		Use:     "down",
		PreRunE: prepMB,
		RunE: func(cmd *cobra.Command, args []string) error {
			return mb.Down(steps)
		},
	}))
	m.AddCommand(&cobra.Command{
		Use:      "reset",
		PreRunE:  prepMB,
		PostRunE: postMigrate,
		RunE: func(cmd *cobra.Command, args []string) error {
			return mb.Reset()
		},
	})
	m.AddCommand(&cobra.Command{
		Use:     "status",
		PreRunE: prepMB,
		RunE: func(cmd *cobra.Command, args []string) error {
			return mb.Status(os.Stdout)
		},
	})

	return m
}

func init() {
	GenerateCmd.AddCommand(MigrateCmd(p.DatabaseConnection, MigrationBox))
	RootCmd.AddCommand(MigrateCmd(p.DatabaseConnection, MigrationBox))
}
