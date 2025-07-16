package database

import (
	"context"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Invoke(RunMigrations),
)

func RunMigrations(lifecycle fx.Lifecycle, db *gorm.DB) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return AutoMigrate(db)
		},
	})
}
