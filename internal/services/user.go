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
)

type UserService struct {
	log     *slog.Logger
	usrRepo UserRepo
}

type UserRepo interface {
	Create(ctx context.Context, usr *entities.UserInfo) (id int, err error)
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
