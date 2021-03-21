package services

import (
	"github.com/duckbrain/shiboleet/lib/config"
	"github.com/duckbrain/shiboleet/models"
	"github.com/gobuffalo/pop/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Config struct {
	config.Config
}

type Provider struct {
	AppName string
	Config
	Repository *models.Repository
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
	p.Repository = &models.Repository{Executor: c.Store.(boil.ContextExecutor)}

	return c.Open()
}
