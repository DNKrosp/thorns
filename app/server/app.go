package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"thorns/health"
	healthHTTP "thorns/health/http"
	"thorns/health/usecase"
	"time"
)

type App struct {
	httpServer *http.Server

	HealthUC health.UseCase
}

func NewApp() *App {
	return &App{
		HealthUC: usecase.NewHealthUseCase(),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	healthHTTP.RegisterHTTPEndpoints(router, a.HealthUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func (s *App) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
