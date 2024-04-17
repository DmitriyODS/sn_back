package pdb

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"idon.com/models"
)

const (
	SqlSelectUser = `
SELECT id
FROM users
WHERE login = $1
  AND password = crypt($2, password);
`
)

func (p *PDB) SelectUser(ctx context.Context, user models.User) (uint64, error) {
	var rows pgx.Rows
	var err error

	rows, err = p.QueryTx(ctx, SqlSelectUser, user.Login, user.Password)
	if rows == nil {
		return 0, models.ErrNoRows
	}
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	userID, err := pgx.CollectOneRow(rows, pgx.RowTo[uint64])
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, models.ErrNoRows
	}
	if err != nil {
		return 0, err
	}

	return userID, nil
}
