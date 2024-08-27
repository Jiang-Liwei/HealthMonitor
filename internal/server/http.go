package server

import (
	helloworldV1 "HealthMonitor/api/helloworld/v1"
	indexV1 "HealthMonitor/api/index/v1"
	"HealthMonitor/internal/biz"
	"HealthMonitor/internal/conf"
	"HealthMonitor/internal/data"
	"HealthMonitor/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, dataData *data.Data, logger log.Logger) *http.Server {
	var serverOpts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		serverOpts = append(serverOpts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		serverOpts = append(serverOpts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		serverOpts = append(serverOpts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(serverOpts...)

	RegisterHTTPRoutes(srv, dataData, logger)
	return srv
}

func RegisterHTTPRoutes(srv *http.Server, dataData *data.Data, logger log.Logger) {
	indexService := service.NewIndexService()
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	helloworldV1.RegisterGreeterHTTPServer(srv, greeterService)
	indexV1.RegisterIndexHTTPServer(srv, indexService)
}
