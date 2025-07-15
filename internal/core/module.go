package core

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
	// fx.Annotate(/*SERVICE_IMPLEMENT, fx.As(new(ports.SERVICE_INTERFACE))*/),
	),
)
