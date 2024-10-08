package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
	"healthmonitor/internal/biz"
)

type adminUserRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdminUserRepo 创建一个新的用户数据存储
func NewAdminUserRepo(data *Data) biz.UserRepo {
	return &adminUserRepo{data: data}
}

// FindByID 根据用户ID查找用户
func (ur *adminUserRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.AdminUser, error) {
	user, err := ur.data.db.AdminUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
