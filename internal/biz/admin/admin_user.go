package admin

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
)

// UserRepo 定义用户存储的接口
type UserRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*ent.AdminUser, error)
	FindByUsername(ctx context.Context, username string) (*ent.AdminUser, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewAdminUserUsecase 创建一个新的后台用户用例
func NewAdminUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// GetUser 获取用户信息
func (uc *UserUsecase) GetUser(ctx context.Context, id uuid.UUID) (*ent.AdminUser, error) {
	return uc.repo.FindByID(ctx, id)
}

// FindByUsername 获取用户信息
func (uc *UserUsecase) FindByUsername(ctx context.Context, username string) (*ent.AdminUser, error) {
	return uc.repo.FindByUsername(ctx, username)
}
