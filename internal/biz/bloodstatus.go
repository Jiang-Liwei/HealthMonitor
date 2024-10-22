package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
	"healthmonitor/ent/bloodstatusrecord"
	"healthmonitor/internal/common"
	"healthmonitor/internal/middleware/auth/jwt"
)

type BloodStatus struct {
	ID                uuid.UUID
	RecordDate        uint
	TimeOfDay         bloodstatusrecord.TimeOfDay
	BeforeAfterMeals  bloodstatusrecord.BeforeAfterMeals
	SystolicPressure  uint8
	DiastolicPressure uint8
	Pulse             uint8
}

// BloodStatusRepo is a Greater repo.
type BloodStatusRepo interface {
	Save(context.Context, *BloodStatus, uuid.UUID) (*ent.BloodStatusRecord, error)
	Update(context.Context, uuid.UUID, *BloodStatus, uuid.UUID) (*ent.BloodStatusRecord, error)
	DeleteBloodStatus(context.Context, uuid.UUID) error
	FindByID(context.Context, uuid.UUID) (*ent.BloodStatusRecord, error)
	ListByUserID(context.Context, uuid.UUID, int, int) (*common.PageData[*ent.BloodStatusRecord], error)
}

// BloodStatusUsecase is a Greeter usecase.
type BloodStatusUsecase struct {
	repo BloodStatusRepo
	log  *log.Helper
}

// NewBloodStatusUsecase new a BloodStatus usecase.
func NewBloodStatusUsecase(repo BloodStatusRepo, logger log.Logger) *BloodStatusUsecase {
	return &BloodStatusUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Create creates a BloodStatus, and returns the new ent.BloodStatusRecord.
func (uc *BloodStatusUsecase) Create(ctx context.Context, b *BloodStatus) (*ent.BloodStatusRecord, error) {
	uc.log.WithContext(ctx).Infof("CreateBloodStatusRecord: %v", b)
	user, _ := jwt.UserFromContext(ctx)
	return uc.repo.Save(ctx, b, user.ID)
}

// Update a BloodStatus, and returns the new ent.BloodStatusRecord.
func (uc *BloodStatusUsecase) Update(ctx context.Context, id uuid.UUID, b *BloodStatus) (*ent.BloodStatusRecord, error) {
	uc.log.WithContext(ctx).Infof("UpdateBloodStatusRecord: %v", b)
	user, _ := jwt.UserFromContext(ctx)
	return uc.repo.Update(ctx, id, b, user.ID)
}

// Delete a BloodStatus.
func (uc *BloodStatusUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	uc.log.WithContext(ctx).Infof("DeleteBloodStatusRecord: %v", id)
	return uc.repo.DeleteBloodStatus(ctx, id)
}

// FindByID returns a ent.BloodStatusRecord.
func (uc *BloodStatusUsecase) FindByID(ctx context.Context, id uuid.UUID) (*ent.BloodStatusRecord, error) {
	return uc.repo.FindByID(ctx, id)
}

// List returns the common.PageData[*ent.BloodStatusRecord].
func (uc *BloodStatusUsecase) List(ctx context.Context, page int, pageSize int) (*common.PageData[*ent.BloodStatusRecord], error) {
	user, _ := jwt.UserFromContext(ctx)
	records, _ := uc.repo.ListByUserID(ctx, user.ID, page, pageSize)
	return records, nil
}
