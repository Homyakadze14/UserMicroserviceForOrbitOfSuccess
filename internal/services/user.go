package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/entities"
	"github.com/google/uuid"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrBadRequest        = errors.New("bad request")
)

type UserService struct {
	log     *slog.Logger
	usrRepo UserRepo
}

type UserRepo interface {
	Create(ctx context.Context, usr *entities.UserInfo) (id int, err error)
	Update(ctx context.Context, usr *entities.UserInfo) error
}

func NewUserService(
	log *slog.Logger,
	usrRepo UserRepo,
) *UserService {
	return &UserService{
		log:     log,
		usrRepo: usrRepo,
	}
}

func (s *UserService) CreateDefault(ctx context.Context, usr *entities.UserInfo) error {
	const op = "User.CreateDefault"

	log := s.log.With(
		slog.String("op", op),
		slog.String("acc", usr.String()),
	)

	log.Info("trying to create default user")
	usr.Firstname = "User-" + uuid.NewString()
	_, err := s.usrRepo.Create(ctx, usr)
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully created default user")

	return nil
}

func (s *UserService) Update(ctx context.Context, usr *entities.UserInfo) error {
	const op = "User.Update"

	log := s.log.With(
		slog.String("op", op),
		slog.String("acc", usr.String()),
	)

	log.Info("trying to update user")
	err := s.usrRepo.Update(ctx, usr)
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully updated user")

	return nil
}
