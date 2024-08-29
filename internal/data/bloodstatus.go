package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"healthmonitor/internal/biz"
)

type bloodStatusRepo struct {
	data *Data
	log  *log.Helper
}

// NewBloodStatusRepo creates a new instance of bloodStatusRepo.
func NewBloodStatusRepo(data *Data, logger log.Logger) biz.BloodStatusRepo {
	return &bloodStatusRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (bsr *bloodStatusRepo) Save(ctx context.Context, b *biz.BloodStatus) (*biz.BloodStatus, error) {
	return b, nil
}

func (bsr *bloodStatusRepo) Update(ctx context.Context, b *biz.BloodStatus) (*biz.BloodStatus, error) {
	return b, nil
}

func (bsr *bloodStatusRepo) FindByID(ctx context.Context, id string) (*biz.BloodStatus, error) {
	return &biz.BloodStatus{}, nil
}

func (bsr *bloodStatusRepo) ListByID(ctx context.Context, id string) ([]*biz.BloodStatus, error) {
	return []*biz.BloodStatus{}, nil
}

func (bsr *bloodStatusRepo) ListAll(ctx context.Context) ([]*biz.BloodStatus, error) {
	return []*biz.BloodStatus{}, nil
}
