package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RestaurantRepository struct {
	db *pgxpool.Pool
}

func NewRestaurantRepository(db *pgxpool.Pool) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

type Restaurant struct {
	ID          string
	Name        string
	Description string
	Address     string
	Phone       string
	Latitude    float64
	Longitude   float64
	LogoURL     string
	OpeningTime string
	ClosingTime string
	CreatedAt   string
	UpdatedAt   string
}

func (r *RestaurantRepository) GetRestaurants(ctx context.Context, offset int32, limit int32) ([]*Restaurant, error) {
	var restaurants []*Restaurant

	query := `
		SELECT id, name, COALESCE(description, ''), address, phone, latitude, longitude, 
		       COALESCE(logo_url, ''), opening_time::text, closing_time::text, 
		       to_char(created_at, 'YYYY-MM-DD HH24:MI:SS'), 
		       to_char(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM restaurants
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query restaurants: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var restaurant Restaurant
		err := rows.Scan(
			&restaurant.ID, &restaurant.Name, &restaurant.Description,
			&restaurant.Address, &restaurant.Phone, &restaurant.Latitude,
			&restaurant.Longitude, &restaurant.LogoURL, &restaurant.OpeningTime,
			&restaurant.ClosingTime, &restaurant.CreatedAt, &restaurant.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan restaurant: %w", err)
		}
		restaurants = append(restaurants, &restaurant)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating restaurants: %w", err)
	}

	return restaurants, nil
}

func (r *RestaurantRepository) GetRestaurantsCount(ctx context.Context) (int32, error) {
	var count int32
	query := `SELECT COUNT(*) FROM restaurants`
	err := r.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count restaurants: %w", err)
	}
	return count, nil
}

func (r *RestaurantRepository) GetRestaurant(ctx context.Context, id string) (*Restaurant, error) {
	var restaurant Restaurant

	query := `
		SELECT id, name, COALESCE(description, ''), address, phone, latitude, longitude, 
		       COALESCE(logo_url, ''), opening_time::text, closing_time::text, 
		       to_char(created_at, 'YYYY-MM-DD HH24:MI:SS'), 
		       to_char(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM restaurants
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&restaurant.ID, &restaurant.Name, &restaurant.Description,
		&restaurant.Address, &restaurant.Phone, &restaurant.Latitude,
		&restaurant.Longitude, &restaurant.LogoURL, &restaurant.OpeningTime,
		&restaurant.ClosingTime, &restaurant.CreatedAt, &restaurant.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get restaurant: %w", err)
	}

	return &restaurant, nil
}

func (r *RestaurantRepository) GetRestaurantStatus(ctx context.Context, restaurantID string) (*Restaurant, error) {
	var restaurant Restaurant

	query := `
		SELECT id, opening_time, closing_time
		FROM restaurants
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, restaurantID).Scan(
		&restaurant.ID,
		&restaurant.OpeningTime, &restaurant.ClosingTime,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get restaurant status: %w", err)
	}

	return &restaurant, nil
}
