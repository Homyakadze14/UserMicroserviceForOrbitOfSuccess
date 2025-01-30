package app

import (
	"fmt"
	"log/slog"
	"os"

	grpcapp "github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/app/grpc"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/config"
	repositories "github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/repositories/postgresql"
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
	usrRepo := repositories.NewUserRepository(pg)

	// Services
	userService := services.NewUserService(log, usrRepo)

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
