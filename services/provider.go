package services

import (
	"github.com/duckbrain/shiboleet/lib/config"
	"github.com/gobuffalo/pop/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Config struct {
	config.Config
}

type Provider struct {
	AppName string
	Config
	DatabaseConnection *pop.Connection
}

func (p Provider) Must() *Provider {
	err := p.Init()
	if err != nil {
		panic(err)
	}
	return &p
}

func (p *Provider) Init() error {
	c, err := pop.NewConnection(p.Database)
	if err != nil {
		return err
	}
	p.DatabaseConnection = c

	return c.Open()
}

func (p *Provider) DB() boil.ContextExecutor {
	if p.DatabaseConnection.TX != nil {
		return p.DatabaseConnection.TX
	}
	return p.DatabaseConnection.Store.(boil.ContextExecutor)
}
