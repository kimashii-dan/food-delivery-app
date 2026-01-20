package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MenuItemRepository struct {
	db *pgxpool.Pool
}

func NewMenuItemRepository(db *pgxpool.Pool) *MenuItemRepository {
	return &MenuItemRepository{db: db}
}

type MenuItem struct {
	ID           string
	RestaurantID string
	Name         string
	Description  string
	Price        float64
	ImageURL     string
	IsAvailable  bool
	Category     string
	CreatedAt    string
	UpdatedAt    string
}

func (r *MenuItemRepository) GetMenu(ctx context.Context, restaurantID string, offset int32, limit int32) ([]*MenuItem, error) {
	var items []*MenuItem

	query := `
		SELECT id, restaurant_id, name, COALESCE(description, ''), price, COALESCE(image_url, ''),
		       is_available, COALESCE(category, ''), 
		       to_char(created_at, 'YYYY-MM-DD HH24:MI:SS'), 
		       to_char(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM menu_items
		WHERE restaurant_id = $1
		ORDER BY category, name
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(ctx, query, restaurantID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query menu items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item MenuItem
		err := rows.Scan(
			&item.ID, &item.RestaurantID, &item.Name, &item.Description,
			&item.Price, &item.ImageURL, &item.IsAvailable, &item.Category,
			&item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan menu item: %w", err)
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating menu items: %w", err)
	}

	return items, nil
}

func (r *MenuItemRepository) GetMenuCount(ctx context.Context, restaurantID string) (int32, error) {
	var count int32
	query := `SELECT COUNT(*) FROM menu_items WHERE restaurant_id = $1`
	err := r.db.QueryRow(ctx, query, restaurantID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count menu items: %w", err)
	}
	return count, nil
}

func (r *MenuItemRepository) GetMenuItem(ctx context.Context, id string) (*MenuItem, error) {
	var item MenuItem

	query := `
		SELECT id, restaurant_id, name, COALESCE(description, ''), price, COALESCE(image_url, ''),
		       is_available, COALESCE(category, ''), 
		       to_char(created_at, 'YYYY-MM-DD HH24:MI:SS'), 
		       to_char(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM menu_items
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&item.ID, &item.RestaurantID, &item.Name, &item.Description,
		&item.Price, &item.ImageURL, &item.IsAvailable, &item.Category,
		&item.CreatedAt, &item.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get menu item: %w", err)
	}

	return &item, nil
}

type MenuItemValidation struct {
	ItemID      string
	IsAvailable bool
	Name        string
}

func (r *MenuItemRepository) ValidateMenuItems(ctx context.Context, restaurantID string, itemIDs []string) ([]*MenuItemValidation, error) {
	var validations []*MenuItemValidation

	query := `
		SELECT id, is_available, name
		FROM menu_items
		WHERE restaurant_id = $1 AND id = ANY($2)
	`

	rows, err := r.db.Query(ctx, query, restaurantID, itemIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to validate menu items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var validation MenuItemValidation
		err := rows.Scan(&validation.ItemID, &validation.IsAvailable, &validation.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan menu item validation: %w", err)
		}
		validations = append(validations, &validation)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating menu item validations: %w", err)
	}

	return validations, nil
}
