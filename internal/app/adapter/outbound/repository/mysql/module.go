package mysql

import (
	"context"

	"github.com/asfung/ticus/internal/core/ports"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Options(
	fx.Provide(
		NewDatabase,
		// fx.Annotate(ports.NewArticleRepository, fx.As(new(ports.ArticleRepository))), // figured out, is this equivalent to IoC world
		ports.NewArticleRepository,
		ports.NewAuthRepository,
	),
	fx.Invoke(RunMigrations),
)

func RunMigrations(lifecycle fx.Lifecycle, db *gorm.DB) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return AutoMigrate(db)
		},
	})
}
