package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/duckbrain/shiboleet/models"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input []*models.User) (*models.CreateUserPayload, error) {
	res := make([]*models.User, 0, len(input))
	for _, m := range input {
		if m.ID == uuid.Nil {
			var err error
			m.ID, err = uuid.NewV4()
			if err != nil {
				return nil, err
			}
		}
		err := m.Insert(ctx, r.DB(), boil.Infer())
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return &models.CreateUserPayload{Items: res}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input []*models.User) (*models.UpdateUserPayload, error) {
	res := make([]*models.User, 0, len(input))
	for _, m := range input {
		_, err := m.Update(ctx, r.DB(), boil.Infer())

		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return &models.UpdateUserPayload{Items: res}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id []string) (*models.DeleteUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "World!", nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return models.Users().All(ctx, r.DB())
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
