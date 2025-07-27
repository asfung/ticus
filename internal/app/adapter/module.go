package adapter

import (
	"github.com/asfung/ticus/internal/app/adapter/inbound"
	"github.com/asfung/ticus/internal/app/adapter/outbound"
	"go.uber.org/fx"
)

var Module = fx.Options(
	inbound.Module,
	outbound.Module,
)
