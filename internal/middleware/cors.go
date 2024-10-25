package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// CORS 中间件，允许跨域请求
func CORS() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			if header, ok := transport.FromServerContext(ctx); ok {
				origin := header.RequestHeader().Get("Origin")
				allowedOrigins := []string{"*"}

				// 检查请求的 Origin 是否被允许
				if isAllowedOrigin(origin, allowedOrigins) {
					println("请求来源被允许: ", origin) // 调试信息
					header.ReplyHeader().Set("Access-Control-Allow-Origin", origin)
					header.ReplyHeader().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
					header.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
					header.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
				}

				// 如果请求是预检请求（OPTIONS），直接返回 200 状态码
				if header.Operation() == http.MethodOptions {
					println("处理 OPTIONS 预检请求") // 调试信息
					// 确保在预检请求中返回正确的跨域头信息
					header.ReplyHeader().Set("Access-Control-Allow-Origin", origin)
					header.ReplyHeader().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
					header.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
					header.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
					header.ReplyHeader().Set("Content-Length", "0")
					header.ReplyHeader().Set("Content-Type", "text/plain; charset=utf-8")
					return nil, nil
				}
			}

			// 继续处理其他请求
			return handler(ctx, req)
		}
	}
}

// isAllowedOrigin 判断请求的 Origin 是否在允许的列表中
func isAllowedOrigin(origin string, allowedOrigins []string) bool {
	if origin == "" {
		return false
	}
	for _, allowedOrigin := range allowedOrigins {
		if strings.EqualFold(origin, allowedOrigin) {
			return true
		}
	}
	return false
}
