package inbound

import (
	"github.com/asfung/ticus/internal/app/adapter/inbound/api"
	"go.uber.org/fx"
)

var Module = fx.Options(
	api.Module,
)
