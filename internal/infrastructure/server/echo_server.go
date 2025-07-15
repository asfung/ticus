package server

import (
	"context"

	"github.com/asfung/ticus/internal/infrastructure/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)


var Module = fx.Options(
	fx.Provide(NewEchoServer),
)


func NewEchoServer() *echo.Echo{
	return echo.New()
}

func StartServer(lifecycle fx.Lifecycle, config *config.AppConfig, e *echo.Echo)  {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func ()  {
				e.Logger.Fatal(e.Start(":" + config.Port))
			}()
			return nil
		},
	})
}



