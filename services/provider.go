package services

import (
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/pop/v5"
)

type Config struct {
	// Environment the application will be running in, eg: "development" and "production"
	Environment string

	// Server options
	ListenAddr string // Address the web server will listen on
	Host       string // Absolute URL the application will be accessed through

	LogLevel logger.Level

	Database *pop.ConnectionDetails

	Assets struct {
		Type                  string
		ManifestPath          string
		DevelopmentServerPath string
	}
}

type Provider struct {
	AppName string
	Config
	DatabaseConnection *pop.Connection
}

func (p Provider) Init() error {
	c, err := pop.NewConnection(p.Database)
	if err != nil {
		return err
	}
	p.DatabaseConnection = c

	return nil
}
