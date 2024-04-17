package httpClient

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"idon.com/models"
)

func (hs *HttpServer) makeErrHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		log.Errorf("Err handler: %v", err)
		resp := models.MakeResponse(fiber.StatusInternalServerError, nil, ErrServer)

		return ctx.Status(resp.Code).JSON(resp)
	}
}
