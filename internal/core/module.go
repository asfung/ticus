package core

import (
	"github.com/asfung/ticus/internal/core/ports"
	"github.com/asfung/ticus/internal/core/services"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	ports.Module,
	fx.Provide(
		func() *logrus.Logger {
			return logrus.New()
		},
		// fx.Annotate(/*SERVICE_IMPLEMENT, fx.As(new(ports.SERVICE_INTERFACE))*/),
		fx.Annotate(services.NewArticleService, fx.As(new(ports.ArticleService))),
	),
)
