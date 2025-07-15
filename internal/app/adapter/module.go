package adapter

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
)
