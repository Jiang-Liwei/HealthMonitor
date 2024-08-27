//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"HealthMonitor/internal/conf"
	"HealthMonitor/internal/server"
	"HealthMonitor/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProvideHTTPServerConfig provides the HTTP server configuration.
func ProvideHTTPServerConfig(
	c *conf.Server,
	index *service.IndexService,
	logger log.Logger,
) server.HTTPServerConfig {
	return server.HTTPServerConfig{
		Conf:   c,
		Index:  index,
		Logger: logger,
	}
}

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		//data.ProviderSet,
		//biz.ProviderSet,
		service.ProviderSet,
		ProvideHTTPServerConfig,
		newApp,
	))
}
