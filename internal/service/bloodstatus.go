package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "healthmonitor/api/bloodstatus/v1"
	"healthmonitor/ent/bloodstatusrecord"
	"healthmonitor/internal/biz"
)

type BloodStatusService struct {
	pb.UnimplementedBloodStatusServer

	uc  *biz.BloodStatusUsecase
	log *log.Helper
}

func NewBloodStatusService(uc *biz.BloodStatusUsecase, logger log.Logger) *BloodStatusService {
	return &BloodStatusService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BloodStatusService) CreateBloodStatus(ctx context.Context, req *pb.CreateBloodStatusRequest) (*pb.CreateBloodStatusReply, error) {
	s.log.Infof("input data %v", req)
	record, err := s.uc.Create(ctx, &biz.BloodStatus{
		RecordDate:        uint(req.RecordDate),
		TimeOfDay:         bloodstatusrecord.TimeOfDay(req.TimeOfDay.String()),
		BeforeAfterMeals:  bloodstatusrecord.BeforeAfterMeals(req.BeforeAfterMeals.String()),
		SystolicPressure:  uint8(req.SystolicPressure),
		DiastolicPressure: uint8(req.DiastolicPressure),
		Pulse:             uint8(req.Pulse),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateBloodStatusReply{
		BloodStatus: &pb.BloodStatusRecord{
			RecordDate:        uint64(record.RecordDate),
			TimeOfDay:         record.TimeOfDay.String(),
			BeforeAfterMeals:  record.BeforeAfterMeals.String(),
			SystolicPressure:  uint32(record.SystolicPressure),
			DiastolicPressure: uint32(record.DiastolicPressure),
			Pulse:             uint32(record.Pulse),
		},
	}, nil
}

func (s *BloodStatusService) UpdateBloodStatus(ctx context.Context, req *pb.UpdateBloodStatusRequest) (*pb.UpdateBloodStatusReply, error) {
	s.log.Infof("input data %v", req)
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	record, err := s.uc.Update(ctx, id, &biz.BloodStatus{
		RecordDate:        uint(req.RecordDate),
		TimeOfDay:         bloodstatusrecord.TimeOfDay(req.TimeOfDay.String()),
		BeforeAfterMeals:  bloodstatusrecord.BeforeAfterMeals(req.BeforeAfterMeals.String()),
		SystolicPressure:  uint8(req.SystolicPressure),
		DiastolicPressure: uint8(req.DiastolicPressure),
		Pulse:             uint8(req.Pulse),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateBloodStatusReply{
		BloodStatus: &pb.BloodStatusRecord{
			RecordDate:        uint64(record.RecordDate),
			TimeOfDay:         record.TimeOfDay.String(),
			BeforeAfterMeals:  record.BeforeAfterMeals.String(),
			SystolicPressure:  uint32(record.SystolicPressure),
			DiastolicPressure: uint32(record.DiastolicPressure),
			Pulse:             uint32(record.Pulse),
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
			RecordDate:        uint64(record.RecordDate),
			TimeOfDay:         record.TimeOfDay.String(),
			BeforeAfterMeals:  record.BeforeAfterMeals.String(),
			SystolicPressure:  uint32(record.SystolicPressure),
			DiastolicPressure: uint32(record.DiastolicPressure),
			Pulse:             uint32(record.Pulse),
		},
	}, nil
}

func (s *BloodStatusService) ListBloodStatus(ctx context.Context, req *pb.ListBloodStatusRequest) (*pb.ListBloodStatusReply, error) {
	s.log.Infof("input data %v", req)

	// 调用 Usecase 层的方法获取数据
	data, err := s.uc.List(ctx, int(req.Page), int(req.PageSize))

	if err != nil {
		return nil, err
	}
	// 转换数据为 Protobuf 格式
	replyData := make([]*pb.BloodStatusRecord, 0)
	for _, v := range data.Data {
		replyData = append(replyData, &pb.BloodStatusRecord{
			RecordDate:        uint64(v.RecordDate),
			TimeOfDay:         v.TimeOfDay.String(),
			BeforeAfterMeals:  v.BeforeAfterMeals.String(),
			SystolicPressure:  uint32(v.SystolicPressure),
			DiastolicPressure: uint32(v.DiastolicPressure),
			Pulse:             uint32(v.Pulse),
		})
	}

	// 构造并返回分页响应
	return &pb.ListBloodStatusReply{
		Data: &pb.PageData{
			Page:       int64(data.Page),
			PageSize:   int64(data.PageSize),
			TotalPages: int64(data.TotalPages),
			TotalCount: int64(data.TotalCount),
			Data:       replyData, // 注意这里使用 replyData
		},
	}, nil
}
