package service

import (
	"github.com/google/wire"
	"healthmonitor/internal/service/admin"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewIndexService,
	NewBloodStatusService,
	admin.NewAuthService,
	admin.NewDashboardService,
)
