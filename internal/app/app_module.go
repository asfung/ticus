package app

import (
	"github.com/asfung/ticus/internal/app/adapter"
	"github.com/asfung/ticus/internal/core"
	"github.com/asfung/ticus/internal/infrastructure/config"
	"github.com/asfung/ticus/internal/infrastructure/database"
	"github.com/asfung/ticus/internal/infrastructure/server"
	"go.uber.org/fx"
)

func NewApp(cfg *config.AppConfig) *fx.App {
	return fx.New(
		fx.Provide(func() *config.AppConfig {
			return cfg
		}),
		database.Module,
		adapter.Module,
		core.Module,
		server.Module,
		fx.Invoke(
			server.StartServer,
		),
	)
}
