package pdb

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"idon.com/models"
)

const (
	SqlSelectPosts = `
SELECT p.id,
       title,
       text,
       p.user_id,
       u.login          as user_name,
       count(l.post_id) as count_likes,
       created_date
FROM posts p
         LEFT JOIN users u on p.user_id = u.id
         LEFT JOIN likes l on p.id = l.post_id
GROUP BY p.id, title, text, p.user_id, u.login, created_date
ORDER BY created_date DESC;
`
	SqlSelectPost = `
SELECT p.id,
       title,
       text,
       p.user_id,
       u.login          as user_name,
       count(l.post_id) as count_likes,
       created_date
FROM posts p
         LEFT JOIN users u on p.user_id = u.id
         LEFT JOIN likes l on p.id = l.post_id
WHERE p.id = $1
GROUP BY p.id, title, text, p.user_id, u.login, created_date
ORDER BY created_date DESC;
`
	SqlInsertPost = `
INSERT INTO posts (title, text, user_id)
VALUES ($1, $2, $3);
`
	SqlUpdatePost = `
UPDATE posts
SET title = $1,
	text = $2
WHERE id = $3;
`
	SqlDeletePost = `
DELETE
FROM posts
WHERE id = $1;
`
)

func (p *PDB) SelectPost(ctx context.Context, id uint64) (models.Post, error) {
	var rows pgx.Rows
	var err error
	var post models.Post

	rows, err = p.QueryTx(ctx, SqlSelectPost, id)
	if rows == nil {
		return post, models.ErrNoRows
	}
	if err != nil {
		return post, err
	}
	defer rows.Close()

	post, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.Post])
	if errors.Is(err, pgx.ErrNoRows) {
		return post, models.ErrNoRows
	}
	if err != nil {
		return post, err
	}

	return post, nil
}

func (p *PDB) SelectPostList(ctx context.Context) (models.Posts, error) {
	var posts models.Posts

	rows, err := p.QueryTx(ctx, SqlSelectPosts)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	posts, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Post])
	if errors.Is(err, pgx.ErrNoRows) {
		return posts, nil
	}
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (p *PDB) InsertPost(ctx context.Context, post models.Post) error {
	commandTag, err := p.ExecTx(ctx, SqlInsertPost, post.Title, post.Text, post.UserID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return models.ErrNoRows
	}

	return nil
}

func (p *PDB) UpdatePost(ctx context.Context, post models.Post) error {
	commandTag, err := p.ExecTx(ctx, SqlUpdatePost, post.Title, post.Text, post.ID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return models.ErrNoRows
	}

	return nil
}

func (p *PDB) DeletePost(ctx context.Context, id uint64) error {
	commandTag, err := p.ExecTx(ctx, SqlDeletePost, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return models.ErrNoRows
	}

	return nil
}
