package biz

import (
	"github.com/google/wire"
	"healthmonitor/internal/biz/admin"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewBloodStatusUsecase,
	admin.NewAdminUserUsecase,
	NewJWTUsecase,
)
