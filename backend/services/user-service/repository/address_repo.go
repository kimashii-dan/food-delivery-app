package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AddressRepository struct {
	db *pgxpool.Pool
}

func NewAddressRepository(db *pgxpool.Pool) *AddressRepository {
	return &AddressRepository{db: db}
}

type Address struct {
	ID         string
	UserID     string
	Street     string
	City       string
	PostalCode string
	Latitude   float64
	Longitude  float64
	IsDefault  bool
	CreatedAt  string
}

func (r *AddressRepository) Create(ctx context.Context, address *Address) error {
	if address.IsDefault {
		updateQuery := `UPDATE addresses SET is_default = false WHERE user_id = $1`
		_, err := r.db.Exec(ctx, updateQuery, address.UserID)
		if err != nil {
			return fmt.Errorf("failed to unset default addresses: %w", err)
		}
	}

	query := `
        INSERT INTO addresses (id, user_id, street, city, postal_code, latitude, longitude, is_default, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
    `

	_, err := r.db.Exec(ctx, query,
		address.ID, address.UserID, address.Street, address.City,
		address.PostalCode, address.Latitude, address.Longitude, address.IsDefault,
	)

	if err != nil {
		return fmt.Errorf("failed to create address: %w", err)
	}

	return nil
}

func (r *AddressRepository) GetByUserID(ctx context.Context, userID string) ([]*Address, error) {
	query := `
        SELECT id, user_id, street, city, postal_code, latitude, longitude, is_default,
               to_char(created_at, 'YYYY-MM-DD HH24:MI:SS')
        FROM addresses 
        WHERE user_id = $1
        ORDER BY is_default DESC, created_at DESC
    `

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query addresses: %w", err)
	}
	defer rows.Close()

	addresses := []*Address{}
	for rows.Next() {
		var addr Address
		err := rows.Scan(
			&addr.ID, &addr.UserID, &addr.Street, &addr.City,
			&addr.PostalCode, &addr.Latitude, &addr.Longitude, &addr.IsDefault,
			&addr.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan address: %w", err)
		}
		addresses = append(addresses, &addr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating addresses: %w", err)
	}

	return addresses, nil
}

func (r *AddressRepository) GetByID(ctx context.Context, addressID string) (*Address, error) {
	var addr Address

	query := `
        SELECT id, user_id, street, city, postal_code, latitude, longitude, is_default
        FROM addresses 
        WHERE id = $1
    `

	err := r.db.QueryRow(ctx, query, addressID).Scan(
		&addr.ID, &addr.UserID, &addr.Street, &addr.City,
		&addr.PostalCode, &addr.Latitude, &addr.Longitude, &addr.IsDefault,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get address: %w", err)
	}

	return &addr, nil
}

func (r *AddressRepository) Delete(ctx context.Context, addressID string) error {
	query := `DELETE FROM addresses WHERE id = $1`

	_, err := r.db.Exec(ctx, query, addressID)
	if err != nil {
		return fmt.Errorf("failed to delete address: %w", err)
	}

	return nil
}
