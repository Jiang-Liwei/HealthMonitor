package admin

import (
	"context"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"healthmonitor/internal/biz"
	"healthmonitor/internal/biz/admin"
	"healthmonitor/internal/middleware/auth"

	pb "healthmonitor/api/admin/v1"
)

type DashboardService struct {
	pb.UnimplementedDashboardServer

	uc  *admin.DashboardUsecase
	log *log.Helper
	jwt *biz.JWTUsecase
}

func NewDashboardService(logger log.Logger, uc *admin.DashboardUsecase, jwt *biz.JWTUsecase) *DashboardService {
	return &DashboardService{
		uc:  uc,
		log: log.NewHelper(logger),
		jwt: jwt,
	}
}

func (s *DashboardService) Dashboard(ctx context.Context, req *pb.DashboardRequest) (*pb.DashboardReply, error) {
	user, ok := auth.UserFromContext(ctx)
	if !ok {
		return nil, errors2.Unauthorized("UNAUTHORIZED", "User is not login")
	}
	reply, err := s.uc.Dashboard(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
