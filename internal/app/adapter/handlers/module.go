package handlers

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api"
	"go.uber.org/fx"
)

var Module = fx.Options(
	api.Module,
)
