package admin

import (
	"context"
	"errors"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "healthmonitor/api/adminauth/v1"
	"healthmonitor/internal/biz"
	"healthmonitor/internal/biz/admin"
	"healthmonitor/internal/middleware/auth"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	uc  *admin.UserUsecase
	log *log.Helper
	jwt *biz.JWTUsecase
}

func NewAuthService(uc *admin.UserUsecase, logger log.Logger, jwt *biz.JWTUsecase) *AuthService {
	return &AuthService{
		uc:  uc,
		log: log.NewHelper(logger),
		jwt: jwt,
	}
}

func (s *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := s.uc.FindByUsername(ctx, req.Username)
	if err != nil || biz.CheckPassword(user.PasswordHash, req.Password) != nil {
		return nil, errors.New("账号或密码错误")
	}
	token, err := s.jwt.GenerateToken(user.ID.String())
	return &pb.SignInResponse{
		AccessToken: token,
		User: &pb.MeResponse{
			Id:          user.ID.String(),
			Username:    user.Username,
			Email:       user.Email,
			IsActive:    user.IsActive,
			LastLoginAt: int64(user.LastLoginAt),
			CreatedAt:   int64(user.CreatedAt),
			UpdatedAt:   int64(user.UpdatedAt),
		},
	}, nil
}

func (s *AuthService) Me(ctx context.Context, req *pb.MeRequest) (*pb.MeResponse, error) {
	user, ok := auth.UserFromContext(ctx)
	if !ok {
		return nil, errors2.Unauthorized("UNAUTHORIZED", "User is not login")
	}

	return &pb.MeResponse{
		Id:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		IsActive:    user.IsActive,
		LastLoginAt: int64(user.LastLoginAt),
		CreatedAt:   int64(user.CreatedAt),
		UpdatedAt:   int64(user.UpdatedAt),
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	claims, ok := auth.TokenFromContext(ctx)
	if !ok {
		return nil, errors2.Unauthorized("UNAUTHORIZED", "无效的令牌")
	}
	token, err := auth.GetStringToken(claims)
	if err != nil {
		return nil, err
	}
	// 将 JWT 加入黑名单
	err = s.jwt.RevokeToken(token)
	if err != nil {
		return nil, errors2.InternalServer("INTERNAL_SERVER_ERROR", "无法将令牌加入黑名单")
	}

	// 返回成功响应
	return &pb.LogoutResponse{
		Code: 200,
		Msg:  "Logout successful",
	}, nil
}
