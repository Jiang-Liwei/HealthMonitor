package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"healthmonitor/ent"
	"healthmonitor/ent/bloodstatusrecord"
	"healthmonitor/internal/common"
	"healthmonitor/internal/middleware/auth"
)

type BloodStatus struct {
	ID                uuid.UUID
	RecordDate        uint
	TimeOfDay         bloodstatusrecord.TimeOfDay
	BeforeAfterMeals  bloodstatusrecord.BeforeAfterMeals
	SystolicPressure  uint8
	DiastolicPressure uint8
	Pulse             uint8
	Mood              bloodstatusrecord.Mood
	StatusSummary     bloodstatusrecord.StatusSummary
}

// BloodStatusRepo is a Greater repo.
type BloodStatusRepo interface {
	Save(context.Context, *BloodStatus, uuid.UUID) (*ent.BloodStatusRecord, error)
	Update(context.Context, uuid.UUID, *BloodStatus, uuid.UUID) (*ent.BloodStatusRecord, error)
	DeleteBloodStatus(context.Context, uuid.UUID) error
	FindByID(context.Context, uuid.UUID) (*ent.BloodStatusRecord, error)
	PageListByUserID(context.Context, uuid.UUID, int, int) (*common.PageData[*ent.BloodStatusRecord], error)
	ListByUserID(context.Context, uuid.UUID, int, int) ([]*ent.BloodStatusRecord, error)
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
	warn := GetStatusSummary(b)
	warn += ""
	uc.log.WithContext(ctx).Infof("CreateBloodStatusRecord: %v", b)
	user, _ := auth.UserFromContext(ctx)
	return uc.repo.Save(ctx, b, user.ID)
}

// Update a BloodStatus, and returns the new ent.BloodStatusRecord.
func (uc *BloodStatusUsecase) Update(ctx context.Context, id uuid.UUID, b *BloodStatus) (*ent.BloodStatusRecord, error) {
	warn := GetStatusSummary(b)
	warn += ""
	uc.log.WithContext(ctx).Infof("UpdateBloodStatusRecord: %v", b)
	user, _ := auth.UserFromContext(ctx)
	return uc.repo.Update(ctx, id, b, user.ID)
}

// Delete a BloodStatus.
func (uc *BloodStatusUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	uc.log.WithContext(ctx).Infof("DeleteBloodStatusRecord: %v", id)
	return uc.repo.DeleteBloodStatus(ctx, id)
}

// FindByID returns a *ent.BloodStatusRecord.
func (uc *BloodStatusUsecase) FindByID(ctx context.Context, id uuid.UUID) (*ent.BloodStatusRecord, error) {
	return uc.repo.FindByID(ctx, id)
}

// List returns the common.PageData[*ent.BloodStatusRecord].
func (uc *BloodStatusUsecase) List(ctx context.Context, start int, end int) ([]*ent.BloodStatusRecord, error) {
	user, _ := auth.UserFromContext(ctx)
	records, _ := uc.repo.ListByUserID(ctx, user.ID, start, end)
	return records, nil
}

func GetStatusSummary(b *BloodStatus) string {
	var score int = 3
	var warn string
	var goodhealth = true
	if b.SystolicPressure < 90 || b.SystolicPressure > 120 {
		goodhealth = false
		warn += " 高压异常"
	}
	if b.DiastolicPressure < 60 || b.DiastolicPressure > 80 {
		goodhealth = false
		warn += " 低压异常"
	}
	if b.Pulse < 60 || b.Pulse > 100 {
		goodhealth = false
		warn += " 脉搏异常"
	}

	if goodhealth {
		score += 1
	} else {
		score -= 1
	}

	switch b.Mood {
	case "happy":
		score += 1
	case "sad":
		score -= 1
		warn += " 不开心咯"
	}

	switch score {
	case 1:
		b.StatusSummary = "very bad"
	case 2:
		b.StatusSummary = "bad"
	case 3:
		b.StatusSummary = "common"
	case 4:
		b.StatusSummary = "good"
	case 5:
		b.StatusSummary = "perfect"
	}

	return warn
}
