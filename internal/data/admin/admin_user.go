package admin

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
	"healthmonitor/ent/adminuser"
	"healthmonitor/internal/biz/admin"
	"healthmonitor/internal/data/core"
)

type adminUserRepo struct {
	data *core.Data
	log  *log.Helper
}

// NewAdminUserRepo 创建一个新的用户数据存储
func NewAdminUserRepo(data *core.Data) admin.UserRepo {
	return &adminUserRepo{data: data}
}

// FindByID 根据用户ID查找用户
func (ur *adminUserRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.AdminUser, error) {
	user, err := ur.data.DB.AdminUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (ur *adminUserRepo) FindByUsername(ctx context.Context, username string) (*ent.AdminUser, error) {
	user, err := ur.data.DB.AdminUser.Query().
		Where(adminuser.UsernameEQ(username)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
