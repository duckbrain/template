// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct {
	Executor boil.ContextExecutor
}

func (r *Repository) DB(ctx context.Context) boil.ContextExecutor {
	return r.Executor
}
