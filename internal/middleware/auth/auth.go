package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v4"
	"healthmonitor/ent"
)

type authKey struct{}

type AuthInfo struct {
	User  *ent.AdminUser
	Token *jwt.Token
}

// NewContext 将认证信息放入上下文
func NewContext(ctx context.Context, info *AuthInfo) context.Context {
	return context.WithValue(ctx, authKey{}, info)
}

// UserFromContext 从上下文中提取 *ent.AdminUser 信息
func UserFromContext(ctx context.Context) (*ent.AdminUser, bool) {
	authInfo, ok := ctx.Value(authKey{}).(*AuthInfo)
	if !ok {
		return nil, false
	}
	return authInfo.User, true
}

// TokenFromContext 从上下文中提取 *jwt.Token 信息
func TokenFromContext(ctx context.Context) (*jwt.Token, bool) {
	authInfo, ok := ctx.Value(authKey{}).(*AuthInfo)
	if !ok {
		return nil, false
	}
	return authInfo.Token, true
}

// FromContext extract auth info from context
func FromContext(ctx context.Context) (*AuthInfo, bool) {
	authInfo, ok := ctx.Value(authKey{}).(*AuthInfo)
	return authInfo, ok
}

func GetStringToken(token *jwt.Token) (string, error) {
	// 提取 JWT ID（jti）
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.InternalServer("INTERNAL_SERVER_ERROR", "无法解析令牌声明")
	}

	jti, ok := claims["jti"].(string)
	if !ok {
		return "", errors.InternalServer("INTERNAL_SERVER_ERROR", "无效的 JWT ID")
	}

	return jti, nil
}
