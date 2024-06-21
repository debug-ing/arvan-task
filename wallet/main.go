package main

import (
	"context"
	"log"
	"net/http"

	"github.com/debug-ing/arvan-task/wallet/api"
	"github.com/debug-ing/arvan-task/wallet/config"
	"github.com/debug-ing/arvan-task/wallet/internal"
	"github.com/debug-ing/arvan-task/wallet/pkg/repository"
	"github.com/debug-ing/arvan-task/wallet/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			NewEcho,
			config.LoadConfig,
			internal.ConnectDB,
			repository.NewWalletRepository,
			repository.NewTransactionRepository,
			service.NewWalletService,
			api.NewWalletRoutes,
		),
		fx.Invoke(api.Register),
		fx.Invoke(StartServer),
	).Run()
}

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("Validator", validator.New())
			return next(c)
		}
	})
	e.Use(internal.CustomLogger())
	return e
}

func StartServer(lc fx.Lifecycle, e *echo.Echo, config *config.AppConfig) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := e.Start(config.Port); err != nil && err != http.ErrServerClosed {
					log.Fatalf("shutting down the server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
