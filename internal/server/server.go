package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"healthmonitor/internal/conf"
	"healthmonitor/internal/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

// HTTPServerConfig holds the configuration and dependencies for HTTP server.
type HTTPServerConfig struct {
	Conf        *conf.Server
	Index       *service.IndexService
	BloodStatus *service.BloodStatusService
	Logger      log.Logger
}

// GRPCServerConfig holds the configuration and dependencies for GRPC server.
type GRPCServerConfig struct {
	Conf   *conf.Server
	Index  *service.IndexService
	Logger log.Logger
}
