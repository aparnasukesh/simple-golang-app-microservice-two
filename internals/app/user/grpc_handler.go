package user

import (
	"context"
	"microservice-two/grpcproto"
)

type GrpcHandler struct {
	grpcproto.UnimplementedMicroServiceTwoServiceServer
	svc Service
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}

func (h *GrpcHandler) MethodOne(ctx context.Context, req *grpcproto.MethodRequest) (*grpcproto.MethodResponse, error) {
	// Sequential processing using service's mutex
	h.svc.MethodOne(int(req.WaitTime))

	// Return the user names received from Microservice-One
	return &grpcproto.MethodResponse{
		UserNames: req.Names,
	}, nil
}

func (h *GrpcHandler) MethodTwo(ctx context.Context, req *grpcproto.MethodRequest) (*grpcproto.MethodResponse, error) {
	// Parallel processing
	h.svc.MethodTwo(int(req.WaitTime))

	// Return the user names received from Microservice-One
	return &grpcproto.MethodResponse{
		UserNames: req.Names,
	}, nil
}
