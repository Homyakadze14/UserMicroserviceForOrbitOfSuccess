package controller

import (
	userv1 "github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/proto/gen/user"
	"google.golang.org/grpc"
)

type serverAPI struct {
	userv1.UnimplementedUserServer
	user User
}

type User interface {
}

func Register(gRPCServer *grpc.Server, user User) {
	userv1.RegisterUserServer(gRPCServer, &serverAPI{user: user})
}
