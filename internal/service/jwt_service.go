package service

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"healthmonitor/ent"
)

type authKey struct{}

type JWTService struct {
	data     *ent.Client
	log      *log.Helper
	secret   []byte
	lifespan time.Duration
}

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

// NewJWTService 创建新的 JWT 服务
func NewJWTService(data *ent.Client, logger log.Logger, secret string, lifespan time.Duration) *JWTService {
	return &JWTService{
		data:     data,
		log:      log.NewHelper(logger),
		secret:   []byte(secret),
		lifespan: lifespan,
	}
}

// GenerateToken 生成 JWT 令牌
func (s *JWTService) GenerateToken(userID string) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.lifespan)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

// ValidateToken 验证 JWT 令牌
func (s *JWTService) ValidateToken(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secret, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// 检查 JWT 是否在黑名单中
	blacklisted, err := s.data.AdminJWTBlacklist.Query().Where( /* 条件，例如 jti=claims.ID */ ).Exist(context.Background())
	if err != nil {
		return "", err
	}

	if blacklisted {
		return "", errors.New("token is blacklisted")
	}

	return claims.UserID, nil
}

// RevokeToken 撤销 JWT 令牌
func (s *JWTService) RevokeToken(jti string) error {
	_, err := s.data.AdminJWTBlacklist.Create().
		SetJti(jti).
		SetExpiresAt(int(time.Now().Add(s.lifespan).Unix())).
		Save(context.Background())

	return err
}
