// server/server.go

package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	bloodStatusV1 "healthmonitor/api/bloodstatus/v1"
	indexV1 "healthmonitor/api/index/v1"
)

// NewHTTPServer creates a new HTTP server with the provided configuration.
func NewHTTPServer(cfg HTTPServerConfig) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if cfg.Conf.Http.Network != "" {
		opts = append(opts, http.Network(cfg.Conf.Http.Network))
	}
	if cfg.Conf.Http.Addr != "" {
		opts = append(opts, http.Address(cfg.Conf.Http.Addr))
	}
	if cfg.Conf.Http.Timeout != nil {
		opts = append(opts, http.Timeout(cfg.Conf.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	indexV1.RegisterIndexHTTPServer(srv, cfg.Index)
	bloodStatusV1.RegisterBloodStatusHTTPServer(srv, cfg.BloodStatus)
	return srv
}
