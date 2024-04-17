package httpClient

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"idon.com/models"
)

func (hs *HttpServer) makeLoginHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var user models.User
		if err := ctx.BodyParser(&user); err != nil {
			return err
		}

		serviceCtx := ctx.Locals(CtxTx).(context.Context)

		respData, err := hs.appService.Login(serviceCtx, user)
		if err != nil {
			return err
		}

		resp := ctx.Locals(CtxResp).(*models.Response)
		resp.Data = respData

		return nil
	}
}
