package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg HTTPServerConfig) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if cfg.Conf.Grpc.Network != "" {
		opts = append(opts, grpc.Network(cfg.Conf.Grpc.Network))
	}
	if cfg.Conf.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(cfg.Conf.Grpc.Addr))
	}
	if cfg.Conf.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(cfg.Conf.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	return srv
}
