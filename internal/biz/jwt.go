package biz

import (
	"context"
	"errors"
	"healthmonitor/ent/adminjwtexpiredtokens"
	"healthmonitor/internal/data/core"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type authKey struct{}

type JWTUsecase struct {
	data     *core.Data
	log      *log.Helper
	secret   []byte
	lifespan time.Duration
}

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

// NewJWTUsecase 创建新的 JWT 服务
func NewJWTUsecase(data *core.Data, logger log.Logger) *JWTUsecase {
	return &JWTUsecase{
		data:     data,
		log:      log.NewHelper(logger),
		secret:   []byte("secret"),
		lifespan: 24 * time.Hour,
	}
}

// GenerateToken 生成 JWT 令牌
func (s *JWTUsecase) GenerateToken(userID string) (string, error) {
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
func (s *JWTUsecase) ValidateToken(tokenString string) (string, error) {
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
	blacklisted, err := s.data.DB.AdminJWTExpiredTokens.Query().Where(adminjwtexpiredtokens.JtiEQ(claims.ID)).Exist(context.Background())
	if err != nil {
		return "", err
	}

	if blacklisted {
		return "", errors.New("token is expire")
	}

	return claims.UserID, nil
}

// RevokeToken 撤销 JWT 令牌
func (s *JWTUsecase) RevokeToken(jti string) error {
	_, err := s.data.DB.AdminJWTExpiredTokens.Create().
		SetJti(jti).
		SetExpiresAt(int(time.Now().Add(s.lifespan).Unix())).
		Save(context.Background())

	return err
}

// KeyFunc 用于返回 JWT 签名密钥
func (s *JWTUsecase) KeyFunc(token *jwt.Token) (interface{}, error) {
	// 检查签名方法是否为 HMAC-SHA256
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}
	// 返回用于签名的密钥
	return s.secret, nil
}
