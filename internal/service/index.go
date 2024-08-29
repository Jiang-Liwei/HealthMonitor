package service

import (
	"context"

	pb "healthmonitor/api/index/v1"
)

type IndexService struct {
	pb.UnimplementedIndexServer
}

func NewIndexService() *IndexService {
	return &IndexService{}
}

func (s *IndexService) Index(ctx context.Context, req *pb.IndexRequest) (*pb.IndexReply, error) {
	return &pb.IndexReply{Message: "哈哈哈哈~成功啦~！"}, nil
}
