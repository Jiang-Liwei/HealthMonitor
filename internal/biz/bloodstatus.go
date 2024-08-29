package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "healthmonitor/api/bloodstatus/v1"
	"time"
)

type BloodStatus struct {
	RecordDate        time.Time
	TimeOfDay         string
	BeforeAfterMeals  string
	SystolicPressure  uint32
	DiastolicPressure uint32
	Pulse             uint32
}

// BloodStatusRepo is a Greater repo.
type BloodStatusRepo interface {
	Save(context.Context, *BloodStatus) (*BloodStatus, error)
	Update(context.Context, *BloodStatus) (*BloodStatus, error)
	FindByID(context.Context, string) (*BloodStatus, error)
	ListByID(context.Context, string) ([]*BloodStatus, error)
	ListAll(context.Context) ([]*BloodStatus, error)
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

// CreateBloodStatus creates a BloodStatus, and returns the new BloodStatus.
func (uc *BloodStatusUsecase) CreateBloodStatus(ctx context.Context, b *pb.CreateBloodStatusRequest) (*BloodStatus, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", b)
	return uc.repo.Save(ctx, nil)
}
