package adapter

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers"
	"github.com/asfung/ticus/internal/app/adapter/repositories"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	repositories.Module,
)
