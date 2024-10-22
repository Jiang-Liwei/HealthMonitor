// server/server.go

package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	adminAuthV1 "healthmonitor/api/adminauth/v1"
	bloodStatusV1 "healthmonitor/api/bloodstatus/v1"
	indexV1 "healthmonitor/api/index/v1"
	"healthmonitor/internal/biz"
	"healthmonitor/internal/biz/admin"
	jwtMiddleware "healthmonitor/internal/middleware/auth/jwt"
)

// NewHTTPServer creates a new HTTP server with the provided configuration.
func NewHTTPServer(cfg HTTPServerConfig, adminUserUsecase *admin.UserUsecase, jwtUsecase *biz.JWTUsecase) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(jwtMiddleware.Server(
				jwtUsecase,
				adminUserUsecase,
			)).Path("/api.adminauth.v1.Auth/User", "/api.adminauth.v1.Auth/Logout").
				Prefix("/api.bloodstatus.v1.BloodStatus/").
				Build(),
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
	opts = append(opts, http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	)))

	srv := http.NewServer(opts...)
	indexV1.RegisterIndexHTTPServer(srv, cfg.Index)
	bloodStatusV1.RegisterBloodStatusHTTPServer(srv, cfg.BloodStatus)
	adminAuthV1.RegisterAuthHTTPServer(srv, cfg.AdminAuth)
	return srv
}
