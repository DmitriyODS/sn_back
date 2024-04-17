package pdb

import (
	"context"
	"idon.com/models"
)

const (
	SqlInsertLike = `
INSERT INTO likes (user_id, post_id)
VALUES ($1, $2);
`
	SqlDeleteLike = `
DELETE
FROM likes
WHERE post_id = $1 AND user_id = $2;
`
)

func (p *PDB) InsertLike(ctx context.Context, like models.Like) error {
	commandTag, err := p.ExecTx(ctx, SqlInsertLike, like.UserID, like.PostID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return models.ErrNoRows
	}

	return nil
}

func (p *PDB) DeleteLike(ctx context.Context, like models.Like) error {
	commandTag, err := p.ExecTx(ctx, SqlDeleteLike, like.UserID, like.PostID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return models.ErrNoRows
	}

	return nil
}
