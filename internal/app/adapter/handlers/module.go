package handlers

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api"
	"github.com/asfung/ticus/internal/app/adapter/handlers/provider"
	"go.uber.org/fx"
)

var Module = fx.Options(
	api.Module,
	provider.Module,
)
