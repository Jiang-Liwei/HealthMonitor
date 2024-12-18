// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"healthmonitor/internal/biz"
	admin2 "healthmonitor/internal/biz/admin"
	"healthmonitor/internal/biz/api"
	"healthmonitor/internal/conf"
	"healthmonitor/internal/data"
	"healthmonitor/internal/data/admin"
	"healthmonitor/internal/server"
	"healthmonitor/internal/service"
	admin3 "healthmonitor/internal/service/admin"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(c *conf.Server, d *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	indexService := service.NewIndexService()
	coreData, cleanup, err := data.NewData(d, logger)
	if err != nil {
		return nil, nil, err
	}
	bloodStatusRepo := data.NewBloodStatusRepo(coreData, logger)
	bloodStatusUsecase := api.NewBloodStatusUsecase(bloodStatusRepo, logger)
	bloodStatusService := service.NewBloodStatusService(bloodStatusUsecase, logger)
	userRepo := admin.NewAdminUserRepo(coreData)
	userUsecase := admin2.NewAdminUserUsecase(userRepo, logger)
	jwtUsecase := biz.NewJWTUsecase(coreData, logger)
	authService := admin3.NewAuthService(userUsecase, logger, jwtUsecase)
	dashboardUsecase := admin2.NewDashboardUsecase(bloodStatusRepo, logger)
	dashboardService := admin3.NewDashboardService(logger, dashboardUsecase, jwtUsecase)
	httpServerConfig := ProvideHTTPServerConfig(c, indexService, bloodStatusService, authService, dashboardService, logger)
	grpcServer := server.NewGRPCServer(httpServerConfig)
	httpServer := server.NewHTTPServer(httpServerConfig, userUsecase, jwtUsecase)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

// ProvideHTTPServerConfig provides the HTTP server configuration.
func ProvideHTTPServerConfig(
	c *conf.Server,
	index *service.IndexService,
	bloodStatus *service.BloodStatusService,
	adminAuth *admin3.AuthService,
	dashboard *admin3.DashboardService,
	logger log.Logger,
) server.HTTPServerConfig {
	return server.HTTPServerConfig{
		Conf:        c,
		Index:       index,
		BloodStatus: bloodStatus,
		AdminAuth:   adminAuth,
		Dashboard:   dashboard,
		Logger:      logger,
	}
}
