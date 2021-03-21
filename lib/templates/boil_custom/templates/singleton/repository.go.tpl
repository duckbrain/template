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
