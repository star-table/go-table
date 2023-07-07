package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/star-table/interface/golang/ping"
)

type PingService struct {
	pb.UnimplementedPingServer
	logger *log.Helper
}

func NewPingService(logger log.Logger) *PingService {
	return &PingService{
		logger: log.NewHelper(logger),
	}
}

func (s *PingService) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{}, nil
}
