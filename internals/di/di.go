package di

import "microservice-two/internals/app/user"

func InitializeResources() (user.GrpcHandler, error) {
	// User Module initialization

	service := user.NewService()
	userGrpcHandler := user.NewGrpcHandler(service)

	return userGrpcHandler, nil
}
