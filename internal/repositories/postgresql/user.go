package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/entities"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/internal/services"
	"github.com/Homyakadze14/UserMicroserviceForOrbitOfSuccess/pkg/postgres"
	"github.com/jackc/pgx/v5"
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
		"INSERT INTO user_info(user_id, firstname, middlename, lastname, gender, phone, icon_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
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

func (r *UserRepository) Update(ctx context.Context, usr *entities.UserInfo) error {
	const op = "repositories.UserRepository.Update"

	_, err := r.Pool.Exec(
		ctx,
		"UPDATE user_info SET firstname=$1, middlename=$2, lastname=$3, gender=$4, phone=$5, icon_url=$6, updated_at=$7 WHERE user_id=$8",
		usr.Firstname, usr.Middlename, usr.Lastname, usr.Gender, usr.Phone, usr.IconURL, time.Now(), usr.UserID)

	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 22001") {
			return services.ErrBadRequest
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *UserRepository) Get(ctx context.Context, uid int) (*entities.UserInfo, error) {
	const op = "repositories.UserRepository.Get"

	row := r.Pool.QueryRow(
		ctx,
		"SELECT firstname, middlename, lastname, gender, phone, icon_url FROM user_info WHERE user_id=$1",
		uid)

	usr := &entities.UserInfo{}
	err := row.Scan(&usr.Firstname, &usr.Middlename, &usr.Lastname, &usr.Gender, &usr.Phone, &usr.IconURL)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return usr, nil
}
