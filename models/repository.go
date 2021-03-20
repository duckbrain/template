package models

import "context"

type Query struct{ Repository }
type Mutation struct{ Repository }

type Repository struct{}

// TODO

func (Repository) CreateUser(ctx context.Context, input []*User) (*CreateUserPayload, error) {
	return nil, nil
}
