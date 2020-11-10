package connectivity

import (
	"context"
	"fghpdf.me/gin-project-template/internal/server/connectivity/pb"
	log "github.com/sirupsen/logrus"
)

type grpcServer struct {
	pb.UnimplementedConnectivityServer
}

func NewGrpcServer() *grpcServer {
	return &grpcServer{}
}

func (s *grpcServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	svc := NewServer()
	res, err := svc.Ping()
	if err != nil {
		log.Errorf("[connectivity][grpcHandler][Ping] svc.Ping error:%v", err)
		return nil, err
	}

	return &pb.PingReply{Status: res.Status}, nil
}
