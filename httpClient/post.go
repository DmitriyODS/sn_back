package httpClient

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"idon.com/models"
)

func (hs *HttpServer) makeGetPostHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.QueryInt("user_id", 0)
		postID, err := ctx.ParamsInt("id", 0)
		if err != nil {
			return err
		}

		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		respData, err := hs.appService.GetPost(serviceCtx, uint64(postID), uint64(userID))
		if err != nil {
			return err
		}

		resp := ctx.Locals(CtxResp).(*models.Response)
		resp.Data = respData

		return nil
	}
}

func (hs *HttpServer) makeGetPostsHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.QueryInt("user_id", 0)
		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		respData, err := hs.appService.GetPosts(serviceCtx, uint64(userID))
		if err != nil {
			return err
		}

		resp := ctx.Locals(CtxResp).(*models.Response)
		resp.Data = respData

		return nil
	}
}

func (hs *HttpServer) makeAddPostHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var post models.Post
		if err := ctx.BodyParser(&post); err != nil {
			return err
		}

		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		if err := hs.appService.AddPost(serviceCtx, post); err != nil {
			return err
		}

		return nil
	}
}

func (hs *HttpServer) makeUpdatePostHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var post models.Post
		if err := ctx.BodyParser(&post); err != nil {
			return err
		}

		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		if err := hs.appService.UpdatePost(serviceCtx, post); err != nil {
			return err
		}

		return nil
	}
}

func (hs *HttpServer) makeDeletePostHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		postID, err := ctx.ParamsInt("id", 0)
		if err != nil {
			return err
		}

		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		if err = hs.appService.DeletePost(serviceCtx, uint64(postID)); err != nil {
			return err
		}

		return nil
	}
}
