package app

import (
	"context"
	"idon.com/models"
)

func (a *App) AddPost(ctx context.Context, post models.Post) error {
	if err := post.Validate(); err != nil {
		return err
	}

	return a.reps.InsertPost(ctx, post)
}

func (a *App) UpdatePost(ctx context.Context, post models.Post) error {
	if err := post.Validate(); err != nil {
		return err
	}

	return a.reps.UpdatePost(ctx, post)
}

func (a *App) DeletePost(ctx context.Context, id uint64) error {
	return a.reps.DeletePost(ctx, id)
}

func (a *App) GetPost(ctx context.Context, id uint64) (models.Post, error) {
	return a.reps.SelectPost(ctx, id)
}

func (a *App) GetPosts(ctx context.Context) (models.Posts, error) {
	return a.reps.SelectPostList(ctx)
}
