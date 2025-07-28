package grpc

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewGRPCServer),
	fx.Invoke(func(server *GRPCServer, lifecycle fx.Lifecycle) {
		server.Start(lifecycle)
	}),
)
