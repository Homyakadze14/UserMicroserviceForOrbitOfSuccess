package controller

import (
	"context"
	"errors"

	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/entities"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/services"
	userv1 "github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/proto/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	userv1.UnimplementedUserServer
	user User
}

type User interface {
	CreateDefault(ctx context.Context, usr *entities.UserInfo) error
	Update(ctx context.Context, usr *entities.UserInfo) error
	Get(ctx context.Context, uid int) (*entities.UserInfo, error)
}

func Register(gRPCServer *grpc.Server, user User) {
	userv1.RegisterUserServer(gRPCServer, &serverAPI{user: user})
}

func (s *serverAPI) CreateDefault(
	ctx context.Context,
	in *userv1.CreateDefaultRequest,
) (*userv1.CreateDefaultResponse, error) {
	data := &entities.UserInfo{
		UserID: int(in.UserId),
	}

	err := s.user.CreateDefault(ctx, data)
	if err != nil {
		if errors.Is(err, services.ErrUserAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to create")
	}

	return &userv1.CreateDefaultResponse{
		Success: true,
	}, nil
}

func (s *serverAPI) UpdateInfo(
	ctx context.Context,
	in *userv1.UpdateInfoRequest,
) (*userv1.UpdateInfoResponse, error) {
	data := &entities.UserInfo{
		UserID:     int(in.UserId),
		Firstname:  in.Firstname,
		Middlename: in.Middlename,
		Lastname:   in.Lastname,
		Gender:     in.Gender,
		Phone:      in.Phone,
		IconURL:    in.IconUrl,
	}

	err := s.user.Update(ctx, data)
	if err != nil {
		if errors.Is(err, services.ErrBadRequest) {
			return nil, status.Error(codes.InvalidArgument, "bad request")
		}
		return nil, status.Error(codes.Internal, "failed to update")
	}

	return &userv1.UpdateInfoResponse{
		Success: true,
	}, nil
}

func (s *serverAPI) GetInfo(
	ctx context.Context,
	in *userv1.GetInfoRequest,
) (*userv1.GetInfoResponse, error) {
	usr, err := s.user.Get(ctx, int(in.UserId))
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get")
	}

	return &userv1.GetInfoResponse{
		Firstname:  usr.Firstname,
		Middlename: usr.Middlename,
		Lastname:   usr.Lastname,
		Gender:     usr.Gender,
		Phone:      usr.Phone,
		IconUrl:    usr.IconURL,
	}, nil
}
