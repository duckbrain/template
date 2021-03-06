package services

import "github.com/gobuffalo/logger"

type Config struct {
	// Environment the application will be running in, eg: "development" and "production"
	Environment string

	// Server options
	ListenAddr string // Address the web server will listen on
	Host       string // Absolute URL the application will be accessed through

	LogLevel logger.Level
}

type Provider struct {
	Config
}
