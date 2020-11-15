package main

import (
	"context"
	"fghpdf.me/gin-project-template/internal/pkg/ctxType"
	"fghpdf.me/gin-project-template/internal/server/connectivity/pb"
	"github.com/google/uuid"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewConnectivityClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx = ctxType.WithRequestIdIntoMetadata(ctx, uuid.New().String())

	r, err := c.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		log.Fatalf("could not check: %v", err)
	}
	log.Printf("Status: %s", r.GetStatus())
}
