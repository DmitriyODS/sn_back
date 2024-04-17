package httpClient

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"idon.com/service"
	"strings"
)

const (
	CtxResp = "response_struct"
	CtxReq  = "request_struct"
	CtxTx   = "tx_struct"
)

var (
	ErrServer     = errors.New("что-то пошло не так")
	ErrRequest    = errors.New("неверные данные в теле запроса")
	ErrAuthorized = errors.New("необходимо авторизироваться")
)

type HttpServer struct {
	appFiber   *fiber.App
	appService service.Service
	appStorage Storage
	rootRouter fiber.Router
}

func MakeHttpServer(srv service.Service, stg Storage) *HttpServer {
	httpServer := &HttpServer{}
	httpServer.InitServer(srv, stg)

	return httpServer
}

func (hs *HttpServer) Listen(addr string) error {
	return hs.appFiber.Listen(addr)
}

func (hs *HttpServer) ShutdownWithContext(ctx context.Context) error {
	return hs.appFiber.ShutdownWithContext(ctx)
}

func (hs *HttpServer) InitServer(srv service.Service, stg Storage) {
	if hs.appFiber != nil {
		return
	}

	hs.appFiber = fiber.New(fiber.Config{
		AppName:           "CRM",
		EnablePrintRoutes: true,
		ErrorHandler:      hs.makeErrHandler(),
	})
	hs.appService = srv
	hs.appStorage = stg

	hs.initFiber()
}

func (hs *HttpServer) initFiber() {
	hs.initGlobalMiddleware()

	hs.rootRouter = hs.appFiber.Group("/api/v1")
	hs.initRootMiddlewares()

	// регистрация маршрутов
	hs.registeredAuthRoutes()
	hs.registerPostRoutes()
	hs.registerLikeRoutes()
}

func (hs *HttpServer) initGlobalMiddleware() {
	// установка ППО CORS
	hs.appFiber.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: strings.Join([]string{
			fiber.HeaderAuthorization,
			fiber.HeaderContentType,
		}, ","),
		ExposeHeaders: "",
		MaxAge:        12,
	}))
}

func (hs *HttpServer) initRootMiddlewares() {
	// регистрация корневого ППО обработки запросов
	hs.rootRouter.Use(hs.makeTransportMiddleware("/api/v1/auth"))

	// регистрация корневого ППО обработки транзакций
	hs.rootRouter.Use(hs.makeTransactionMiddleware())
}

func (hs *HttpServer) registeredAuthRoutes() {
	auth := hs.rootRouter.Group("/auth")

	auth.Post("/login", hs.makeLoginHandler())
}

func (hs *HttpServer) registerPostRoutes() {
	posts := hs.rootRouter.Group("/posts")

	posts.Get("/", hs.makeGetPostsHandler())
	posts.Get("/:id", hs.makeGetPostHandler())

	posts.Post("/", hs.makeAddPostHandler())
	posts.Put("/", hs.makeUpdatePostHandler())
	posts.Delete("/:id", hs.makeDeletePostHandler())
}

func (hs *HttpServer) registerLikeRoutes() {
	likes := hs.rootRouter.Group("/likes")

	likes.Post("/", hs.makeAddLikeHandler())
	likes.Delete("/", hs.makeDeleteLikeHandler())
}
