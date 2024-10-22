package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
	"healthmonitor/ent/bloodstatusrecord"
	"healthmonitor/internal/biz"
	"healthmonitor/internal/common"
	"healthmonitor/internal/data/core"
)

type bloodStatusRepo struct {
	data *core.Data
	log  *log.Helper
}

// NewBloodStatusRepo creates a new instance of bloodStatusRepo.
func NewBloodStatusRepo(data *core.Data, logger log.Logger) biz.BloodStatusRepo {
	return &bloodStatusRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (bsr *bloodStatusRepo) Save(ctx context.Context, b *biz.BloodStatus, userID uuid.UUID) (*ent.BloodStatusRecord, error) {
	record, err := bsr.data.DB.BloodStatusRecord.Create().
		SetUserID(userID).
		SetRecordDate(b.RecordDate).
		SetTimeOfDay(b.TimeOfDay).
		SetBeforeAfterMeals(b.BeforeAfterMeals).
		SetSystolicPressure(b.SystolicPressure).
		SetDiastolicPressure(b.DiastolicPressure).
		SetPulse(b.Pulse).
		Save(ctx)
	return record, err
}

func (bsr *bloodStatusRepo) Update(ctx context.Context, id uuid.UUID, b *biz.BloodStatus, userID uuid.UUID) (*ent.BloodStatusRecord, error) {

	p, err := bsr.data.DB.BloodStatusRecord.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if p.UserID != userID {
		return nil, errors.New("user id not match")
	}

	record, err := p.Update().
		SetRecordDate(b.RecordDate).
		SetTimeOfDay(b.TimeOfDay).
		SetBeforeAfterMeals(b.BeforeAfterMeals).
		SetSystolicPressure(b.SystolicPressure).
		SetDiastolicPressure(b.DiastolicPressure).
		SetPulse(b.Pulse).
		Save(ctx)
	return record, nil
}

func (bsr *bloodStatusRepo) DeleteBloodStatus(ctx context.Context, id uuid.UUID) error {
	return bsr.data.DB.BloodStatusRecord.DeleteOneID(id).Exec(ctx)
}

func (bsr *bloodStatusRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.BloodStatusRecord, error) {
	p, err := bsr.data.DB.BloodStatusRecord.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (bsr *bloodStatusRepo) ListByUserID(ctx context.Context, userID uuid.UUID, page int, pageSize int) (*common.PageData[*ent.BloodStatusRecord], error) {

	// 获取总记录数
	totalCount, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(bloodstatusrecord.UserIDEQ(userID)).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	// 获取实际数据
	records, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(bloodstatusrecord.UserIDEQ(userID)).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order(ent.Desc("record_date")).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return common.NewPageData(page, pageSize, totalCount, records), nil
}
