package httpClient

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"idon.com/cfg"
	"idon.com/models"
	"idon.com/utils"
	"strings"
)

func (hs *HttpServer) authorizedUser(ctx *fiber.Ctx) error {
	req := ctx.Locals(CtxReq).(*models.ClientData)

	accessToken := ctx.Get(fiber.HeaderAuthorization, "")
	if accessToken == "" {
		return ErrAuthorized
	}

	// функция возвращает подпись ключа
	keyFunc := func(t *jwt.Token) (any, error) {
		return []byte(cfg.GetAppConfig().SecretKey), nil
	}

	claims, err := utils.ParseAndValidateJWT(accessToken, keyFunc)

	// проверяем доступ
	if err != nil {
		log.Errorf("Error in authorizedUser: %v", err)
		return ErrAuthorized
	}

	req.UserID = utils.ConvStringToUint64(claims.ID)

	return nil
}

// makeTransportMiddleware - корневая ППО для принятия запросов и отправки ответов
func (hs *HttpServer) makeTransportMiddleware(skipUrls ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		defer func() {
			resp, ok := ctx.Locals(CtxResp).(*models.Response)
			if !ok {
				resp = models.MakeResponse(fiber.StatusInternalServerError, nil, ErrServer)
			}

			if err != nil {
				resp.SetResponse(fiber.StatusInternalServerError, nil, err)
				err = nil
			}

			// отлавливаем панику, если была
			if r := recover(); r != nil {
				log.Errorf("Panic handler: %v", r)
				resp.SetResponse(fiber.StatusInternalServerError, nil, ErrServer)
			}

			err = ctx.Status(resp.Code).JSON(resp)
		}()

		req := &models.ClientData{}
		resp := &models.Response{
			Ok:   true,
			Code: 200,
		}
		ctx.Locals(CtxReq, req)
		ctx.Locals(CtxResp, resp)

		curPath := ctx.Path()

		// проверяем URLs, которым не нужна автризация
		for _, checkUrl := range skipUrls {
			if strings.HasPrefix(curPath, checkUrl) {
				return ctx.Next()
			}
		}

		// пытаемся авторизовать пользователя
		err = hs.authorizedUser(ctx)
		if err != nil {
			resp.SetResponse(fiber.StatusUnauthorized, nil, err)
			return nil
		}

		err = ctx.Next()
		if err != nil {
			log.Errorf("Error in handler: %v", err)
			return err
		}

		return nil
	}
}

func (hs *HttpServer) makeTransactionMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		newCtx, tx := hs.appStorage.ContextWithTx(context.Background())

		if tx == nil {
			return nil
		}

		defer func() {
			if txErr := tx.Rollback(newCtx); txErr != nil && !errors.Is(txErr, pgx.ErrTxClosed) {
				log.Errorf("Err rollback transaction: %v", txErr)
				if err == nil {
					err = txErr
				}
			}
		}()

		ctx.Locals(CtxTx, newCtx)

		err = ctx.Next()
		if err != nil {
			log.Errorf("Error in handler: %v", err)
			return err
		}

		if err = tx.Commit(newCtx); err != nil {
			return err
		}

		return nil
	}
}
