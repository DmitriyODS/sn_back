package app

import (
	"context"
	"fmt"
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

func (a *App) GetPost(ctx context.Context, post_id uint64, user_id uint64) (models.Post, error) {
	fmt.Println(post_id, user_id)
	return a.reps.SelectPost(ctx, post_id, user_id)
}

func (a *App) GetPosts(ctx context.Context, user_id uint64) (models.Posts, error) {
	return a.reps.SelectPostList(ctx, user_id)
}
