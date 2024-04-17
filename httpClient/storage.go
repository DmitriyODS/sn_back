package httpClient

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Storage interface {
	ContextWithTx(ctx context.Context) (context.Context, pgx.Tx)
}
