package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Name         string
	Phone        string
	Role         string
}

func (r *UserRepository) Create(ctx context.Context, user *User) error {
	query := `
        INSERT INTO users (id, email, password_hash, name, phone, role, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
    `

	_, err := r.db.Exec(ctx, query,
		user.ID, user.Email, user.PasswordHash,
		user.Name, user.Phone, user.Role,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*User, error) {
	var user User

	query := `
        SELECT id, email, password_hash, name, phone, role 
        FROM users 
        WHERE id = $1
    `

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash,
		&user.Name, &user.Phone, &user.Role,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	query := `
        SELECT id, email, password_hash, name, phone, role 
        FROM users 
        WHERE email = $1
    `

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash,
		&user.Name, &user.Phone, &user.Role,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var count int64

	query := `
	SELECT COUNT(*)
	FROM users 
	WHERE email = $1`

	err := r.db.QueryRow(ctx, query, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}
	return count > 0, nil
}

func (r *UserRepository) Update(ctx context.Context, userID, name, phone string) error {
	query := `
		UPDATE users 
		SET name = $1, phone = $2, updated_at = NOW()
		WHERE id = $3
	`

	_, err := r.db.Exec(ctx, query, name, phone, userID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
