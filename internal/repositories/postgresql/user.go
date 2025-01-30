package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/entities"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/services"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/pkg/postgres"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

func (r *UserRepository) Create(ctx context.Context, usr *entities.UserInfo) (id int, err error) {
	const op = "repositories.UserRepository.Create"

	row := r.Pool.QueryRow(
		ctx,
		"INSERT INTO user_info(user_id, firstname, middlename, lastname, gender, phone, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		usr.UserID, usr.Firstname, usr.Middlename, usr.Lastname, usr.Gender, usr.Phone, usr.IconURL, time.Now(), time.Now())

	err = row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return -1, services.ErrUserAlreadyExists
		}
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
