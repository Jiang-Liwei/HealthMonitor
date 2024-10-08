package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
)

// UserRepo 定义用户存储的接口
type UserRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*ent.AdminUser, error)
}

type AdminUserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewAdminUserUsecase 创建一个新的后台用户用例
func NewAdminUserUsecase(repo UserRepo, logger log.Logger) *AdminUserUsecase {
	return &AdminUserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// GetUser 获取用户信息
func (uc *AdminUserUsecase) GetUser(ctx context.Context, id uuid.UUID) (*ent.AdminUser, error) {
	return uc.repo.FindByID(ctx, id)
}
