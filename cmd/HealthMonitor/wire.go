//go:build wireinject
// +build wireinject

package main

import (
	"healthmonitor/internal/biz"
	"healthmonitor/internal/conf"
	"healthmonitor/internal/data"
	"healthmonitor/internal/server"
	"healthmonitor/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProvideHTTPServerConfig provides the HTTP server configuration.
func ProvideHTTPServerConfig(
	c *conf.Server,
	index *service.IndexService,
	bloodStatus *service.BloodStatusService,
	logger log.Logger,
) server.HTTPServerConfig {
	return server.HTTPServerConfig{
		Conf:        c,
		Index:       index,
		BloodStatus: bloodStatus,
		Logger:      logger,
	}
}

// wireApp init kratos application.
func wireApp(c *conf.Server, d *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		ProvideHTTPServerConfig,
		newApp,
	))
}
