// server/server.go

package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	adminV1 "healthmonitor/api/admin/v1"
	adminAuthV1 "healthmonitor/api/adminauth/v1"
	bloodStatusV1 "healthmonitor/api/bloodstatus/v1"
	indexV1 "healthmonitor/api/index/v1"
	"healthmonitor/internal/biz"
	"healthmonitor/internal/biz/admin"
	"healthmonitor/internal/middleware"
	jwtMiddleware "healthmonitor/internal/middleware/auth/jwt"
)

// NewHTTPServer creates a new HTTP server with the provided configuration.
func NewHTTPServer(cfg HTTPServerConfig, adminUserUsecase *admin.UserUsecase, jwtUsecase *biz.JWTUsecase) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.CORS(), // 注册 CORS 中间件  但是没有用。。。。
			selector.Server(jwtMiddleware.Server(
				jwtUsecase,
				adminUserUsecase,
			)).Path("/api.adminauth.v1.Auth/Me", "/api.adminauth.v1.Auth/Logout").
				Prefix("/api.bloodstatus.v1.BloodStatus/", "/api.admin.v1.Dashboard/").
				Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
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
	adminAuthV1.RegisterAuthHTTPServer(srv, cfg.AdminAuth)
	adminV1.RegisterDashboardHTTPServer(srv, cfg.Dashboard)
	return srv
}
