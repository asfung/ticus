package provider

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/provider/oauth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// jwtProvider.Module("JUST-LIKE-THIS-FOR-RIGHT-NOW"),
	oauth.Module,
)
