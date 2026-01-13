package repository

import (
	"alfdwirhmn/bioskop/internal/data/entity"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserRepository interface {
	Register(ctx context.Context, user *entity.Users) (*entity.Users, error)
	FindByIdentifier(ctx context.Context, identifier string) (*entity.Users, error)

	IsEmailExists(ctx context.Context, email string, excludeID int) (bool, error)
	IsUsernameExists(ctx context.Context, username string, excludeID int) (bool, error)
}

type userRepo struct {
	DB   *pgxpool.Pool
	Logg *zap.Logger
}

func NewUserRepository(db *pgxpool.Pool, log *zap.Logger) UserRepository {
	return &userRepo{
		DB:   db,
		Logg: log,
	}
}

func (ur *userRepo) Register(ctx context.Context, user *entity.Users) (*entity.Users, error) {
	query := `
		INSERT INTO users (username, email, password, full_name, phone_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, email, password, full_name, phone_number, created_at, updated_at
	`

	var usr entity.Users
	err :=
		ur.DB.QueryRow(
			ctx,
			query,
			user.Username,
			user.Email,
			user.Password,
			user.FullName,
			user.PhoneNumber,
		).Scan(
			&usr.ID,
			&usr.Username,
			&usr.Email,
			&usr.Password,
			&usr.FullName,
			&usr.PhoneNumber,
			&usr.CreatedAt,
			&usr.UpdatedAt,
		)

	if err != nil {
		ur.Logg.Error("Failed to register", zap.Error(err))
		return nil, err
	}

	ur.Logg.Info("Registered succesfully")
	return &usr, nil
}

func (r *userRepo) FindByIdentifier(ctx context.Context, identifier string) (*entity.Users, error) {
	query := `
		SELECT id, username, email, password, full_name, phone_number, created_at, updated_at
		FROM users
		WHERE email = $1 OR username = $1
	`

	var user entity.Users
	err := r.DB.QueryRow(
		ctx,
		query,
		identifier,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepo) IsEmailExists(ctx context.Context, email string, excludeID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND id != $2)`

	var exists bool
	err := ur.DB.QueryRow(
		ctx,
		query,
		email,
		excludeID,
	).Scan(&exists)
	if err != nil {
		ur.Logg.Error("Failed to check email existence", zap.Error(err))
		return false, err
	}

	return exists, nil
}

func (ur *userRepo) IsUsernameExists(ctx context.Context, username string, excludeID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 AND id != $2)`

	var exists bool
	err := ur.DB.QueryRow(
		ctx,
		query,
		username,
		excludeID,
	).Scan(&exists)
	if err != nil {
		ur.Logg.Error("Failed to check username existence", zap.Error(err))
		return false, err
	}

	return exists, nil
}
