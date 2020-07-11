package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "sample-grpc-service/proto"

	"google.golang.org/grpc"
)

type sampleGrpcServer struct {
	pb.UnimplementedSampleServer
}

func (s *sampleGrpcServer) Running(ctx context.Context, request *pb.RunningRequest) (*pb.RunningResponse, error) {
	fmt.Println("service is running")
	return &pb.RunningResponse{Running: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSampleServer(grpcServer, &sampleGrpcServer{})
	grpcServer.Serve(lis)
}
