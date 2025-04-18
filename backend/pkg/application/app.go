package application

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qed-tech/egts-pkg-debugger/internal/di"
)

type Application struct {
	Container *di.Container
}

func NewApplication(container *di.Container) *Application {
	return &Application{
		Container: container,
	}
}

func (a *Application) Run() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/api/v1/egts/decode", a.Container.Adapters.EGTS.Decode)
	e.GET("/api/v1/egts/history", a.Container.Adapters.EGTS.GetHistory)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(":9090"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
