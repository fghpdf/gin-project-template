package main

import (
	"fghpdf.me/gin-project-template/internal/server/connectivity"
	"fghpdf.me/gin-project-template/internal/server/connectivity/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	connectivitySrv := connectivity.NewGrpcServer()
	pb.RegisterConnectivityServer(s, connectivitySrv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
