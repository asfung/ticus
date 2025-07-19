package api

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/handler"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(handler.NewCommonHandler),
	fx.Provide(handler.NewArticleHandler),
	fx.Provide(handler.NewAuthHandler),
	fx.Provide(handler.NewOAuthHandler),

	fx.Provide(NewRouter),

	fx.Invoke(BasicMiddleware),
	fx.Invoke(RegisterRoutes),
)
