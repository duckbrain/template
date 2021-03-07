package db

import (
	"os"

	"github.com/duckbrain/shiboleet/services"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/cobra"
)

func MigrateCmd(p *services.Provider, box packd.Box) *cobra.Command {
	m := &cobra.Command{
		Use: "migrate",
	}

	var mb pop.MigrationBox

	prepMB := func(cmd *cobra.Command, args []string) (err error) {
		mb, err = pop.NewMigrationBox(box, p.DatabaseConnection)
		return
	}

	var steps int

	withSteps := func(c *cobra.Command) *cobra.Command {
		c.Flags().IntVarP(&steps, "steps", "s", -1, "steps to move")
		return c
	}

	m.AddCommand(withSteps(&cobra.Command{
		Use:     "up",
		PreRunE: prepMB,
		RunE: func(cmd *cobra.Command, args []string) error {
			return mb.Up()
		},
	}))
	m.AddCommand(withSteps(&cobra.Command{
		Use:     "down",
		PreRunE: prepMB,
		RunE: func(cmd *cobra.Command, args []string) error {
			return mb.Down(1)
		},
	}))
	m.AddCommand(&cobra.Command{
		Use:     "reset",
		PreRunE: prepMB,
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
	m.AddCommand(&cobra.Command{
		Use:     "init",
		PreRunE: prepMB,
		RunE: func(cmd *cobra.Command, args []string) error {
			return mb.CreateSchemaMigrations()
		},
	})

	return m
}
