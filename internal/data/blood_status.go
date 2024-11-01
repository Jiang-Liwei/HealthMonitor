package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
	"healthmonitor/ent/bloodstatusrecord"
	"healthmonitor/internal/biz/api"
	"healthmonitor/internal/common"
	"healthmonitor/internal/data/core"
	"time"
)

type bloodStatusRepo struct {
	data *core.Data
	log  *log.Helper
}

// NewBloodStatusRepo creates a new instance of bloodStatusRepo.
func NewBloodStatusRepo(data *core.Data, logger log.Logger) api.BloodStatusRepo {
	return &bloodStatusRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (bsr *bloodStatusRepo) Save(ctx context.Context, b *api.BloodStatus, userID uuid.UUID) (*ent.BloodStatusRecord, error) {
	timeInt := time.Now().UnixMilli()
	record, err := bsr.data.DB.BloodStatusRecord.Create().
		SetUserID(userID).
		SetRecordDate(b.RecordDate).
		SetTimeOfDay(b.TimeOfDay).
		SetBeforeAfterMeals(b.BeforeAfterMeals).
		SetSystolicPressure(b.SystolicPressure).
		SetDiastolicPressure(b.DiastolicPressure).
		SetPulse(b.Pulse).
		SetMood(b.Mood).
		SetStatusSummary(b.StatusSummary).
		SetCreatedAt(timeInt).
		SetUpdatedAt(timeInt).
		Save(ctx)
	return record, err
}

func (bsr *bloodStatusRepo) Update(ctx context.Context, id uuid.UUID, b *api.BloodStatus, userID uuid.UUID) (*ent.BloodStatusRecord, error) {

	p, err := bsr.data.DB.BloodStatusRecord.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	timeInt := time.Now().UnixMilli()
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
		SetMood(b.Mood).
		SetStatusSummary(b.StatusSummary).
		SetUpdatedAt(timeInt).
		Save(ctx)
	return record, nil
}

func (bsr *bloodStatusRepo) DeleteBloodStatus(ctx context.Context, id uuid.UUID) error {
	return bsr.data.DB.BloodStatusRecord.
		UpdateOneID(id).
		SetDeletedAt(time.Now().UnixMilli()).
		Exec(ctx)
}

func (bsr *bloodStatusRepo) FindByID(ctx context.Context, id uuid.UUID) (*ent.BloodStatusRecord, error) {
	p, err := bsr.data.DB.BloodStatusRecord.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (bsr *bloodStatusRepo) PageListByUserID(ctx context.Context, userID uuid.UUID, page int, pageSize int) (*common.PageData[*ent.BloodStatusRecord], error) {

	// 获取总记录数
	totalCount, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(
			bloodstatusrecord.UserIDEQ(userID),
			bloodstatusrecord.DeletedAt(0),
		).Count(ctx)
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

func (bsr *bloodStatusRepo) ListByUserID(ctx context.Context, userID uuid.UUID, start int, end int) ([]*ent.BloodStatusRecord, error) {

	// 获取实际数据
	records, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(
			bloodstatusrecord.UserIDEQ(userID),
			bloodstatusrecord.DeletedAt(0),
		).Order(ent.Desc("record_date")).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (bsr *bloodStatusRepo) CountByUserID(ctx context.Context, userID uuid.UUID) (int, error) {
	num, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(
			bloodstatusrecord.UserIDEQ(userID),
			bloodstatusrecord.DeletedAt(0),
		).Count(ctx)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (bsr *bloodStatusRepo) CountByStatusSummaryAndUserID(ctx context.Context, userID uuid.UUID, statusSummary bloodstatusrecord.StatusSummary) (int, error) {
	num, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(
			bloodstatusrecord.UserIDEQ(userID),
			bloodstatusrecord.DeletedAt(0),
			bloodstatusrecord.StatusSummaryEQ(statusSummary),
		).Count(ctx)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (bsr *bloodStatusRepo) GetPeriodNum(ctx context.Context, userID uuid.UUID, startTime int64, endTime int64) (int, error) {
	num, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(
			bloodstatusrecord.UserIDEQ(userID),
			bloodstatusrecord.DeletedAt(0),
			bloodstatusrecord.CreatedAtGTE(startTime),
			bloodstatusrecord.CreatedAtLTE(endTime),
		).Count(ctx)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (bsr *bloodStatusRepo) GetByStatusSummaryAndPeriodNum(ctx context.Context, userID uuid.UUID, startTime int64, endTime int64, statusSummary bloodstatusrecord.StatusSummary) (int, error) {
	num, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(
			bloodstatusrecord.UserIDEQ(userID),
			bloodstatusrecord.DeletedAt(0),
			bloodstatusrecord.StatusSummaryEQ(statusSummary),
			bloodstatusrecord.CreatedAtGTE(startTime),
			bloodstatusrecord.CreatedAtLTE(endTime),
		).Count(ctx)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (bsr *bloodStatusRepo) GetRecentRecords(ctx context.Context, userID uuid.UUID, limit int) ([]*ent.BloodStatusRecord, error) {
	records, err := bsr.data.DB.BloodStatusRecord.Query().
		Where(
			bloodstatusrecord.UserIDEQ(userID),
			bloodstatusrecord.DeletedAt(0),
		).
		Order(ent.Desc(bloodstatusrecord.FieldCreatedAt)).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return records, nil
}
