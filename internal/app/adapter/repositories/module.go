package repositories

import (
	"github.com/asfung/ticus/internal/app/adapter/repositories/mysql"
	"go.uber.org/fx"
)

var Module = fx.Options(
	mysql.Module,
)
