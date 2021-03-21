package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/duckbrain/shiboleet/models"
	"github.com/gofrs/uuid"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input []*models.User) (*models.CreateUserPayload, error) {
	return r.Repository.CreateUsers(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input []*models.User) (*models.UpdateUserPayload, error) {
	return r.Repository.UpdateUsers(ctx, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id []uuid.UUID) (*models.DeleteUserPayload, error) {
	return r.Repository.DeleteUsers(ctx, id)
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "World!", nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.Repository.Users(ctx, models.UserFilter{})
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
