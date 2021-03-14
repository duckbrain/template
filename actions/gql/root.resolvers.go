package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/duckbrain/shiboleet/models"
)

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "World!", nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return models.Users().All(ctx, r.DB())
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
