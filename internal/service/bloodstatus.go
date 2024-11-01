package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "healthmonitor/api/bloodstatus/v1"
	"healthmonitor/ent/bloodstatusrecord"
	"healthmonitor/internal/biz/api"
)

type BloodStatusService struct {
	pb.UnimplementedBloodStatusServer

	uc  *api.BloodStatusUsecase
	log *log.Helper
}

func NewBloodStatusService(uc *api.BloodStatusUsecase, logger log.Logger) *BloodStatusService {
	return &BloodStatusService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BloodStatusService) CreateBloodStatus(ctx context.Context, req *pb.CreateBloodStatusRequest) (*pb.CreateBloodStatusReply, error) {
	s.log.Infof("input data %v", req)
	record, err := s.uc.Create(ctx, &api.BloodStatus{
		RecordDate:        uint(req.RecordDate),
		TimeOfDay:         bloodstatusrecord.TimeOfDay(req.TimeOfDay.String()),
		BeforeAfterMeals:  bloodstatusrecord.BeforeAfterMeals(req.BeforeAfterMeals.String()),
		SystolicPressure:  uint8(req.SystolicPressure),
		DiastolicPressure: uint8(req.DiastolicPressure),
		Pulse:             uint8(req.Pulse),
		Mood:              bloodstatusrecord.Mood(req.Mood.String()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateBloodStatusReply{
		BloodStatus: &pb.BloodStatusRecord{
			Id:                record.ID.String(),
			RecordDate:        uint64(record.RecordDate),
			TimeOfDay:         record.TimeOfDay.String(),
			BeforeAfterMeals:  record.BeforeAfterMeals.String(),
			SystolicPressure:  uint32(record.SystolicPressure),
			DiastolicPressure: uint32(record.DiastolicPressure),
			Pulse:             uint32(record.Pulse),
			Mood:              string(record.Mood),
			StatusSummary:     string(record.StatusSummary),
		},
	}, nil
}

func (s *BloodStatusService) UpdateBloodStatus(ctx context.Context, req *pb.UpdateBloodStatusRequest) (*pb.UpdateBloodStatusReply, error) {
	s.log.Infof("input data %v", req)
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	record, err := s.uc.Update(ctx, id, &api.BloodStatus{
		RecordDate:        uint(req.RecordDate),
		TimeOfDay:         bloodstatusrecord.TimeOfDay(req.TimeOfDay.String()),
		BeforeAfterMeals:  bloodstatusrecord.BeforeAfterMeals(req.BeforeAfterMeals.String()),
		SystolicPressure:  uint8(req.SystolicPressure),
		DiastolicPressure: uint8(req.DiastolicPressure),
		Pulse:             uint8(req.Pulse),
		Mood:              bloodstatusrecord.Mood(req.Mood.String()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateBloodStatusReply{
		BloodStatus: &pb.BloodStatusRecord{
			Id:                record.ID.String(),
			RecordDate:        uint64(record.RecordDate),
			TimeOfDay:         record.TimeOfDay.String(),
			BeforeAfterMeals:  record.BeforeAfterMeals.String(),
			SystolicPressure:  uint32(record.SystolicPressure),
			DiastolicPressure: uint32(record.DiastolicPressure),
			Pulse:             uint32(record.Pulse),
			Mood:              string(record.Mood),
			StatusSummary:     string(record.StatusSummary),
		},
	}, nil
}

func (s *BloodStatusService) DeleteBloodStatus(ctx context.Context, req *pb.DeleteBloodStatusRequest) (*pb.DeleteBloodStatusReply, error) {
	s.log.Infof("input data %v", req)
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = s.uc.Delete(ctx, id)
	return &pb.DeleteBloodStatusReply{}, err
}

func (s *BloodStatusService) GetBloodStatus(ctx context.Context, req *pb.GetBloodStatusRequest) (*pb.GetBloodStatusReply, error) {
	s.log.Infof("input data %v", req)
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	record, err := s.uc.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.GetBloodStatusReply{
		BloodStatus: &pb.BloodStatusRecord{
			Id:                record.ID.String(),
			RecordDate:        uint64(record.RecordDate),
			TimeOfDay:         record.TimeOfDay.String(),
			BeforeAfterMeals:  record.BeforeAfterMeals.String(),
			SystolicPressure:  uint32(record.SystolicPressure),
			DiastolicPressure: uint32(record.DiastolicPressure),
			Pulse:             uint32(record.Pulse),
			Mood:              string(record.Mood),
			StatusSummary:     string(record.StatusSummary),
		},
	}, nil
}

func (s *BloodStatusService) ListBloodStatus(ctx context.Context, req *pb.ListBloodStatusRequest) (*pb.ListBloodStatusReply, error) {
	s.log.Infof("input data %v", req)

	// 调用 Usecase 层的方法获取数据
	data, err := s.uc.List(ctx, int(req.StartTime), int(req.EndTime))

	if err != nil {
		return nil, err
	}
	// 转换数据为 Protobuf 格式
	replyData := make([]*pb.BloodStatusRecord, 0)
	for _, v := range data {
		replyData = append(replyData, &pb.BloodStatusRecord{
			Id:                v.ID.String(),
			RecordDate:        uint64(v.RecordDate),
			TimeOfDay:         v.TimeOfDay.String(),
			BeforeAfterMeals:  v.BeforeAfterMeals.String(),
			SystolicPressure:  uint32(v.SystolicPressure),
			DiastolicPressure: uint32(v.DiastolicPressure),
			Pulse:             uint32(v.Pulse),
			Mood:              string(v.Mood),
			StatusSummary:     string(v.StatusSummary),
		})
	}

	// 构造并返回分页响应
	return &pb.ListBloodStatusReply{
		Data: replyData,
	}, nil
}
