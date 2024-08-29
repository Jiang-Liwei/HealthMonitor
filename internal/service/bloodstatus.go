package service

import (
	"context"
	"healthmonitor/internal/biz"

	pb "healthmonitor/api/bloodstatus/v1"
)

type BloodStatusService struct {
	pb.UnimplementedBloodStatusServer
	uc *biz.BloodStatusUsecase
}

func NewBloodStatusService(uc *biz.BloodStatusUsecase) *BloodStatusService {
	return &BloodStatusService{uc: uc}
}

func (s *BloodStatusService) CreateBloodStatus(ctx context.Context, req *pb.CreateBloodStatusRequest) (*pb.CreateBloodStatusReply, error) {
	_, err := s.uc.CreateBloodStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBloodStatusReply{}, nil
}
func (s *BloodStatusService) UpdateBloodStatus(ctx context.Context, req *pb.UpdateBloodStatusRequest) (*pb.UpdateBloodStatusReply, error) {
	return &pb.UpdateBloodStatusReply{}, nil
}
func (s *BloodStatusService) DeleteBloodStatus(ctx context.Context, req *pb.DeleteBloodStatusRequest) (*pb.DeleteBloodStatusReply, error) {
	return &pb.DeleteBloodStatusReply{}, nil
}
func (s *BloodStatusService) GetBloodStatus(ctx context.Context, req *pb.GetBloodStatusRequest) (*pb.GetBloodStatusReply, error) {
	return &pb.GetBloodStatusReply{}, nil
}
func (s *BloodStatusService) ListBloodStatus(ctx context.Context, req *pb.ListBloodStatusRequest) (*pb.ListBloodStatusReply, error) {
	return &pb.ListBloodStatusReply{}, nil
}
