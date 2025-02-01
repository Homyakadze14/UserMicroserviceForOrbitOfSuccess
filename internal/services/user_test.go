package services

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/entities"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type cfg struct {
	usrRepo *mocks.UserRepo
}

func NewService(cfg cfg) *UserService {
	ctx := context.Background()

	usrRepo := cfg.usrRepo
	if cfg.usrRepo == nil {
		usrRepo := &mocks.UserRepo{}
		usrRepo.On("Create", ctx, mock.AnythingOfType("*entities.UserInfo")).Return(0, nil).Once()
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return NewUserService(log, usrRepo)
}

func TestCreateDefault(t *testing.T) {
	// Service config
	ctx := context.Background()
	usrRepo := &mocks.UserRepo{}
	usrRepo.On("Create", ctx, mock.AnythingOfType("*entities.UserInfo")).Return(0, nil).Once()
	sCfg := cfg{
		usrRepo: usrRepo,
	}

	// Test data
	usr := &entities.UserInfo{
		UserID: 1,
	}

	// Function
	service := NewService(sCfg)
	err := service.CreateDefault(ctx, usr)

	// Check
	assert.Nil(t, err)
}

func TestCreateDefaultErr(t *testing.T) {
	// Err
	testErr := errors.New("test")

	// Service config
	ctx := context.Background()
	usrRepo := &mocks.UserRepo{}
	usrRepo.On("Create", ctx, mock.AnythingOfType("*entities.UserInfo")).Return(0, testErr).Once()
	sCfg := cfg{
		usrRepo: usrRepo,
	}

	// Test data
	usr := &entities.UserInfo{
		UserID: 1,
	}

	// Function
	service := NewService(sCfg)
	err := service.CreateDefault(ctx, usr)

	// Check
	assert.ErrorIs(t, err, testErr)
}

func TestUpdate(t *testing.T) {
	// Test data
	usr := &entities.UserInfo{
		UserID:     1,
		Firstname:  "test",
		Middlename: "test",
		Lastname:   "test",
		Phone:      "test",
		Gender:     "m",
	}

	// Service config
	ctx := context.Background()
	usrRepo := &mocks.UserRepo{}
	usrRepo.On("Update", ctx, usr).Return(nil).Once()
	sCfg := cfg{
		usrRepo: usrRepo,
	}

	// Function
	service := NewService(sCfg)
	err := service.Update(ctx, usr)

	// Check
	assert.Nil(t, err)
}

func TestUpdateErr(t *testing.T) {
	// Test data
	usr := &entities.UserInfo{
		UserID:     1,
		Firstname:  "test",
		Middlename: "test",
		Lastname:   "test",
		Phone:      "test",
		Gender:     "m",
	}

	// Err
	testErr := errors.New("test")

	// Service config
	ctx := context.Background()
	usrRepo := &mocks.UserRepo{}
	usrRepo.On("Update", ctx, usr).Return(testErr).Once()
	sCfg := cfg{
		usrRepo: usrRepo,
	}

	// Function
	service := NewService(sCfg)
	err := service.Update(ctx, usr)

	// Check
	assert.ErrorIs(t, err, testErr)
}

func TestGet(t *testing.T) {
	// Test data
	uid := 1
	tUsr := &entities.UserInfo{}

	// Service config
	ctx := context.Background()
	usrRepo := &mocks.UserRepo{}
	usrRepo.On("Get", ctx, uid).Return(tUsr, nil).Once()
	sCfg := cfg{
		usrRepo: usrRepo,
	}

	// Function
	service := NewService(sCfg)
	usr, err := service.Get(ctx, uid)

	// Check
	assert.Nil(t, err)
	assert.Equal(t, tUsr, usr)
}

func TestGetErr(t *testing.T) {
	// Test data
	uid := 1

	// Err
	testErr := errors.New("test")

	// Service config
	ctx := context.Background()
	usrRepo := &mocks.UserRepo{}
	usrRepo.On("Get", ctx, uid).Return(nil, testErr).Once()
	sCfg := cfg{
		usrRepo: usrRepo,
	}

	// Function
	service := NewService(sCfg)
	usr, err := service.Get(ctx, uid)

	// Check
	assert.ErrorIs(t, err, testErr)
	assert.Nil(t, usr)
}
