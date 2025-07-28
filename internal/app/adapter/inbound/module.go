package inbound

import (
	"github.com/asfung/ticus/internal/app/adapter/inbound/api"
	"github.com/asfung/ticus/internal/app/adapter/inbound/grpc"
	"go.uber.org/fx"
)

var Module = fx.Options(
	api.Module,
	grpc.Module,
)
