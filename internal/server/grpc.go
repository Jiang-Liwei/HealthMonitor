package server

import (
	v1 "HealthMonitor/api/index/v1"
	"HealthMonitor/internal/conf"
	"HealthMonitor/internal/data"
	"HealthMonitor/internal/service"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type GRPCServerOptions struct {
	Network      string
	Addr         string
	Timeout      time.Duration
	IndexService *service.IndexService
	Logger       log.Logger
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, dataData *data.Data, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	RegisterGRPCRoutes(srv, dataData)
	return srv
}

func RegisterGRPCRoutes(srv *grpc.Server, dataData *data.Data) {
	indexService := service.NewIndexService()
	v1.RegisterIndexServer(srv, indexService)
}
