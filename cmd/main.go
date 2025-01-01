package main

import (
	"log"
	"microservice-two/config"
	"microservice-two/grpcproto"
	"microservice-two/internals/app/user"
	"microservice-two/internals/di"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	userGrpcHandler, err := di.InitializeResources()
	if err != nil {
		log.Fatalf("error during resource initialization: %v", err)
	}

	server, err := NewGrpcServer(cfg, userGrpcHandler)
	if err != nil {
		log.Fatalf("failed to create gRPC server: %v", err)
	}

	if err := server(); err != nil {
		log.Fatalf("gRPC server encountered an error: %v", err)
	}
}

func NewGrpcServer(cfg config.Config, userGrpcHandler user.GrpcHandler) (func() error, error) {
	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	grpcproto.RegisterMicroServiceTwoServiceServer(s, &userGrpcHandler)
	srv := func() error {
		log.Printf("gRPC server started on port %s", cfg.GrpcPort)
		if err := s.Serve(lis); err != nil {
			log.Printf("failed to serve gRPC: %v", err)
			return err
		}
		return nil
	}
	return srv, nil
}
