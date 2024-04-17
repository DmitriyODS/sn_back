package httpClient

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"idon.com/models"
)

func (hs *HttpServer) makeToggleLikeHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var like models.Like
		if err := ctx.BodyParser(&like); err != nil {
			return err
		}

		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		if err := hs.appService.ToggleLike(serviceCtx, like); err != nil {
			return err
		}

		return nil
	}
}

func (hs *HttpServer) makeAddLikeHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var like models.Like
		if err := ctx.BodyParser(&like); err != nil {
			return err
		}

		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		if err := hs.appService.AddLike(serviceCtx, like); err != nil {
			return err
		}

		return nil
	}
}

func (hs *HttpServer) makeDeleteLikeHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var like models.Like
		if err := ctx.BodyParser(&like); err != nil {
			return err
		}

		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		if err := hs.appService.DeleteLike(serviceCtx, like); err != nil {
			return err
		}

		return nil
	}
}
