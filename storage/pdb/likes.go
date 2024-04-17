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
WHERE user_id = $1 AND post_id = $2;
`
	SqlToggleLike = `
SELECT toggle_like($1, $2);
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

func (p *PDB) ToggleLike(ctx context.Context, like models.Like) error {
	commandTag, err := p.ExecTx(ctx, SqlToggleLike, like.UserID, like.PostID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return models.ErrNoRows
	}

	return nil
}
