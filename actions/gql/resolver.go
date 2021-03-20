package gql

import (
	"github.com/duckbrain/shiboleet/models"
	"github.com/duckbrain/shiboleet/services"
)

type Resolver struct {
	*services.Provider
	models.Repository
}
