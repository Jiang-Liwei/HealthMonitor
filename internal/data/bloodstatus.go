package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
	"healthmonitor/ent/bloodstatusrecord"
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

const UserIDUUID = "550e8400-e29b-41d4-a716-446655440000"

func (bsr *bloodStatusRepo) Save(ctx context.Context, b *biz.BloodStatus) (*ent.BloodStatusRecord, error) {
	userID, err := uuid.Parse(UserIDUUID)
	if err != nil {
		return nil, err
	}
	record, err := bsr.data.db.BloodStatusRecord.Create().
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

func (bsr *bloodStatusRepo) Update(ctx context.Context, id uuid.UUID, b *biz.BloodStatus) (*ent.BloodStatusRecord, error) {

	p, err := bsr.data.db.BloodStatusRecord.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(UserIDUUID)
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
	return bsr.data.db.BloodStatusRecord.DeleteOneID(id).Exec(ctx)
}

func (bsr *bloodStatusRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.BloodStatusRecord, error) {
	p, err := bsr.data.db.BloodStatusRecord.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (bsr *bloodStatusRepo) ListByUserID(ctx context.Context, page int, pageSize int) (*PageData[*ent.BloodStatusRecord], error) {

	userID, err := uuid.Parse(UserIDUUID)
	if err != nil {
		return nil, err
	}

	// 获取总记录数
	totalCount, err := bsr.data.db.BloodStatusRecord.Query().
		Where(bloodstatusrecord.UserIDEQ(userID)).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	// 获取实际数据
	records, err := bsr.data.db.BloodStatusRecord.Query().
		Where(bloodstatusrecord.UserIDEQ(userID)).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order(ent.Desc("record_date")).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return NewPageData[*ent.BloodStatusRecord](page, pageSize, totalCount, records), nil
}
