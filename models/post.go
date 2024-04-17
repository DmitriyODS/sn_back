package models

import (
	"context"
	"errors"
	"time"
)

type Post struct {
	ID              uint64    `json:"id"`
	Title           string    `json:"title"`
	Text            string    `json:"text"`
	UserID          uint64    `json:"user_id"`
	CountLikes      int       `json:"count_likes"`
	CreatedDateUnix int64     `json:"created_date"`
	CreatedDate     time.Time `json:"-"`
}

type Posts []Post

type PostRepository interface {
	SelectPosts(ctx context.Context) (Posts, error)
	SelectPost(ctx context.Context, id uint64) (Post, error)
	InsertPost(ctx context.Context, post Post) error
	UpdatePost(ctx context.Context, post Post) error
	DeletePost(ctx context.Context, id uint64) error
}

func (p Post) Validate() error {
	if p.Title == "" {
		return errors.New("заголовок не может быть пустым")
	}

	return nil
}
