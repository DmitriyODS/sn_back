package models

import (
	"context"
	"errors"
	"time"
)

type Post struct {
	ID              uint64    `json:"id" db:"id"`
	Title           string    `json:"title" db:"title"`
	Text            string    `json:"text" db:"text"`
	UserID          uint64    `json:"user_id" db:"user_id"`
	UserName        string    `json:"user_name" db:"user_name"`
	CountLikes      int       `json:"count_likes" db:"-"`
	CreatedDateUnix int64     `json:"created_date" db:"-"`
	CreatedDate     time.Time `json:"-" db:"created_date"`
}

type Posts []Post

type PostRepository interface {
	SelectPostList(ctx context.Context) (Posts, error)
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
