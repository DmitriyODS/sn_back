package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"idon.com/cfg"
	"idon.com/httpClient"
	"idon.com/service/app"
	"idon.com/storage/pdb"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// инициализируем конфиг
	appConfig := cfg.GetAppConfig()

	// создаём основной контекст
	ctx := context.Background()

	// создаём хранилище
	stg := pdb.GetPDB(ctx)
	defer pdb.ClosePDB()

	// создаём главный сервис
	crmApp := app.MakeApp(stg)

	// создаём сервер
	httpServer := httpClient.MakeHttpServer(crmApp, stg)

	// создаём слушатель сигналов, которые хотят нас прикрыть
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL)
	defer signal.Stop(quit)

	// запускаем горутину с нашим сервером
	go func() {
		log.Info("Server is running...")

		addr := fmt.Sprintf("%s:%s", appConfig.AppAddr, appConfig.AppPort)
		log.Info(addr)
		if err := httpServer.Listen(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("Err in server: %s\n", err)
		}
	}()

	// ждём сигналов о закрытии
	<-quit

	log.Info("Server is shutting down...")
	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := httpServer.ShutdownWithContext(newCtx); err != nil {
		log.Errorf("Err in server shutdown: %s\n", err)
	}
}
