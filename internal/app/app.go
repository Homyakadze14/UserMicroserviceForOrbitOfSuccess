package app

import (
	"fmt"
	"log/slog"
	"os"

	grpcapp "github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/app/grpc"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/config"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/services"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/pkg/postgres"
)

type App struct {
	db         *postgres.Postgres
	GRPCServer *grpcapp.App
}

func Run(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	// Database
	pg, err := postgres.New(cfg.Database.URL, postgres.MaxPoolSize(cfg.Database.PoolMax))
	if err != nil {
		slog.Error(fmt.Errorf("app - Run - postgres.New: %w", err).Error())
		os.Exit(1)
	}

	// Repository

	// Services
	userService := services.NewUserService(log)

	// GRPC
	gRPCServer := grpcapp.New(log, userService, cfg.GRPC.Port)

	return &App{
		db:         pg,
		GRPCServer: gRPCServer,
	}
}

func (s *App) Shutdown() {
	defer s.db.Close()
	defer s.GRPCServer.Stop()
}
