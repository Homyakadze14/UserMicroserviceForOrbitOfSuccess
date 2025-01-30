package services

import (
	"log/slog"
)

type UserService struct {
	log *slog.Logger
}

func NewUserService(
	log *slog.Logger,
) *UserService {
	return &UserService{
		log: log,
	}
}
