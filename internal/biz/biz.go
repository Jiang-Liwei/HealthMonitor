package biz

import (
	"github.com/google/wire"
	"healthmonitor/internal/biz/admin"
	"healthmonitor/internal/biz/api"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	api.NewBloodStatusUsecase,
	admin.NewAdminUserUsecase,
	NewJWTUsecase,
	admin.NewDashboardUsecase,
)
