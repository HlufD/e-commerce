package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/HlufD/users-ms/internals/domain"
)

type UserRepositoryAdapter struct {
	db *sql.DB
}

func NewUserRepositoryAdapter(db *sql.DB) *UserRepositoryAdapter {
	return &UserRepositoryAdapter{
		db,
	}
}

func (ur *UserRepositoryAdapter) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, email, password, created_at, updated_at
	`
	row := ur.db.QueryRowContext(ctx, query,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)

	// Scan the result into the user struct
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepositoryAdapter) FindById(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = $1`

	row := ur.db.QueryRowContext(ctx, query, id)

	// Scan the result into a user struct
	var user domain.User

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepositoryAdapter) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE username = $1`

	row := ur.db.QueryRowContext(ctx, query, username)

	// Scan the result into a user struct
	var user domain.User

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepositoryAdapter) CheckIfUserExists(ctx context.Context, filterKey, filterValue string) (bool, error) {
	allowedKeys := map[string]bool{
		"email":    true,
		"username": true,
		"id":       true,
	}

	if !allowedKeys[filterKey] {
		return false, fmt.Errorf("invalid key for filter")
	}

	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM users WHERE %s = $1)`, filterKey)

	var exists bool

	err := ur.db.QueryRowContext(ctx, query, filterValue).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
