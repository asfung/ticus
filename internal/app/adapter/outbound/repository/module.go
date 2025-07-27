package repository

import (
	"github.com/asfung/ticus/internal/app/adapter/outbound/repository/mysql"
	"go.uber.org/fx"
)

var Module = fx.Options(
	mysql.Module,
)
