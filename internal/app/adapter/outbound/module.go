package outbound

import (
	"github.com/asfung/ticus/internal/app/adapter/outbound/oauth"
	"github.com/asfung/ticus/internal/app/adapter/outbound/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// jwtProvider.Module("JUST-LIKE-THIS-FOR-RIGHT-NOW"),
	oauth.Module,
	repository.Module,
)
