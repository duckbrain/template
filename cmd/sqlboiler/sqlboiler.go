package main

import (
	"os"
	"path/filepath"

	"github.com/duckbrain/shiboleet/lib/config"
	"github.com/duckbrain/shiboleet/lib/runner"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boilingcore"
	"github.com/volatiletech/sqlboiler/v4/drivers"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
	"github.com/volatiletech/sqlboiler/v4/importers"
	"github.com/volatiletech/sqlboiler/v4/templatebin"
)

const AppName = "shiboleet"

var conf = config.Default(AppName)

var defaultTemplatesOut = "lib/templates/boil_default"

func makeTemplates() error {
	_, err := os.Stat(defaultTemplatesOut)
	if !os.IsNotExist(err) {
		return err
	}
	for _, name := range templatebin.AssetNames() {
		p := filepath.Join(defaultTemplatesOut, name)
		err := os.MkdirAll(filepath.Dir(p), os.ModePerm)
		if err != nil {
			return err
		}
		err = os.WriteFile(p, templatebin.MustAsset(name), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

var BoilCmd = &cobra.Command{
	Use: "boil",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return config.Load(AppName, &conf)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := makeTemplates(); err != nil {
			return err
		}
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
			Wipe:      true,
			Imports:   importers.NewDefaultImports(),
			TypeReplaces: []boilingcore.TypeReplace{
				{
					Match:   drivers.Column{DBType: "uuid"},
					Replace: drivers.Column{Type: "uuid.UUID"},
					Imports: importers.Set{ThirdParty: importers.List{`"github.com/gofrs/uuid"`}},
				},
			},
			TemplateDirs: []string{
				filepath.Join(defaultTemplatesOut, "templates"),
				filepath.Join(defaultTemplatesOut, "templates_test"),
				"lib/templates/boil_custom/templates",
			},
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
