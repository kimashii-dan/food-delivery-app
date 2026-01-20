package service

import (
	"context"
	"fmt"
	"time"

	"github.com/kimashii-dan/food-delivery-app/backend/services/restaurant-service/pb"
	"github.com/kimashii-dan/food-delivery-app/backend/services/restaurant-service/repository"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantServiceServer
	restaurantRepo *repository.RestaurantRepository
	menuItemRepo   *repository.MenuItemRepository
}

func NewRestaurantService(restaurantRepo *repository.RestaurantRepository, menuItemRepo *repository.MenuItemRepository) *RestaurantService {
	return &RestaurantService{
		restaurantRepo: restaurantRepo,
		menuItemRepo:   menuItemRepo,
	}
}

func (s *RestaurantService) GetRestaurants(ctx context.Context, req *pb.GetRestaurantsRequest) (*pb.GetRestaurantsResponse, error) {
	page := req.Page
	if page < 1 {
		page = 1
	}

	var limit int32 = 10
	offset := (page - 1) * limit

	restaurants, err := s.restaurantRepo.GetRestaurants(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get restaurants: %v", err)
	}

	total, err := s.restaurantRepo.GetRestaurantsCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get restaurants count: %v", err)
	}

	pbRestaurants := make([]*pb.Restaurant, len(restaurants))
	for i, r := range restaurants {
		pbRestaurants[i] = &pb.Restaurant{
			Id:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Address:     r.Address,
			Phone:       r.Phone,
			Latitude:    r.Latitude,
			Longitude:   r.Longitude,
			LogoUrl:     r.LogoURL,
			OpeningTime: r.OpeningTime,
			ClosingTime: r.ClosingTime,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		}
	}

	return &pb.GetRestaurantsResponse{
		Restaurants: pbRestaurants,
		Total:       total,
	}, nil
}

func (s *RestaurantService) GetRestaurant(ctx context.Context, req *pb.GetRestaurantRequest) (*pb.GetRestaurantResponse, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("restaurant id is required")
	}

	restaurant, err := s.restaurantRepo.GetRestaurant(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("restaurant not found: %v", err)
	}

	return &pb.GetRestaurantResponse{
		Restaurant: &pb.Restaurant{
			Id:          restaurant.ID,
			Name:        restaurant.Name,
			Description: restaurant.Description,
			Address:     restaurant.Address,
			Phone:       restaurant.Phone,
			Latitude:    restaurant.Latitude,
			Longitude:   restaurant.Longitude,
			LogoUrl:     restaurant.LogoURL,
			OpeningTime: restaurant.OpeningTime,
			ClosingTime: restaurant.ClosingTime,
			CreatedAt:   restaurant.CreatedAt,
			UpdatedAt:   restaurant.UpdatedAt,
		},
	}, nil
}

func (s *RestaurantService) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	if req.RestaurantId == "" {
		return nil, fmt.Errorf("restaurant id is required")
	}

	page := req.Page
	if page < 1 {
		page = 1
	}

	var limit int32 = 10
	offset := (page - 1) * limit

	items, err := s.menuItemRepo.GetMenu(ctx, req.RestaurantId, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get menu items: %v", err)
	}

	total, err := s.menuItemRepo.GetMenuCount(ctx, req.RestaurantId)
	if err != nil {
		return nil, fmt.Errorf("failed to get menu items count: %v", err)
	}

	pbItems := make([]*pb.MenuItem, len(items))
	for i, item := range items {
		pbItems[i] = &pb.MenuItem{
			Id:           item.ID,
			RestaurantId: item.RestaurantID,
			Name:         item.Name,
			Description:  item.Description,
			Price:        item.Price,
			ImageUrl:     item.ImageURL,
			IsAvailable:  item.IsAvailable,
			Category:     item.Category,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
	}

	return &pb.GetMenuResponse{
		Items: pbItems,
		Total: total,
	}, nil
}

func (s *RestaurantService) GetMenuItem(ctx context.Context, req *pb.GetMenuItemRequest) (*pb.GetMenuItemResponse, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("menu item id is required")
	}

	item, err := s.menuItemRepo.GetMenuItem(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("menu item not found: %v", err)
	}

	return &pb.GetMenuItemResponse{
		Item: &pb.MenuItem{
			Id:           item.ID,
			RestaurantId: item.RestaurantID,
			Name:         item.Name,
			Description:  item.Description,
			Price:        item.Price,
			ImageUrl:     item.ImageURL,
			IsAvailable:  item.IsAvailable,
			Category:     item.Category,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		},
	}, nil
}

func (s *RestaurantService) GetRestaurantStatus(ctx context.Context, req *pb.GetRestaurantStatusRequest) (*pb.GetRestaurantStatusResponse, error) {
	if req.RestaurantId == "" {
		return nil, fmt.Errorf("restaurant id is required")
	}

	restaurant, err := s.restaurantRepo.GetRestaurantStatus(ctx, req.RestaurantId)
	if err != nil {
		return nil, fmt.Errorf("restaurant not found: %v", err)
	}

	isAcceptingOrders := s.isWithinOpeningHours(restaurant.OpeningTime, restaurant.ClosingTime)

	return &pb.GetRestaurantStatusResponse{
		IsAcceptingOrders: isAcceptingOrders,
		OpeningTime:       restaurant.OpeningTime,
		ClosingTime:       restaurant.ClosingTime,
	}, nil
}

func (s *RestaurantService) ValidateMenuItems(ctx context.Context, req *pb.ValidateMenuItemsRequest) (*pb.ValidateMenuItemsResponse, error) {
	if req.RestaurantId == "" {
		return nil, fmt.Errorf("restaurant id is required")
	}

	if len(req.ItemIds) == 0 {
		return nil, fmt.Errorf("at least one item id is required")
	}

	// Get validations from repository
	validations, err := s.menuItemRepo.ValidateMenuItems(ctx, req.RestaurantId, req.ItemIds)
	if err != nil {
		return nil, fmt.Errorf("failed to validate menu items: %v", err)
	}

	// Convert to map for easy lookup
	validationMap := make(map[string]*repository.MenuItemValidation)
	for _, v := range validations {
		validationMap[v.ItemID] = v
	}

	// Build response
	pbValidations := make([]*pb.MenuItemValidation, 0, len(req.ItemIds))
	unavailableItems := make([]string, 0)
	allAvailable := true

	for _, itemID := range req.ItemIds {
		validation, exists := validationMap[itemID]

		if !exists {
			// Item doesn't exist in this restaurant
			allAvailable = false
			unavailableItems = append(unavailableItems, itemID)
			pbValidations = append(pbValidations, &pb.MenuItemValidation{
				ItemId:      itemID,
				IsAvailable: false,
				Name:        "",
			})
		} else {
			// Item exists, check availability
			if !validation.IsAvailable {
				allAvailable = false
				unavailableItems = append(unavailableItems, itemID)
			}
			pbValidations = append(pbValidations, &pb.MenuItemValidation{
				ItemId:      validation.ItemID,
				IsAvailable: validation.IsAvailable,
				Name:        validation.Name,
			})
		}
	}

	return &pb.ValidateMenuItemsResponse{
		AllAvailable:     allAvailable,
		Items:            pbValidations,
		UnavailableItems: unavailableItems,
	}, nil
}

func (s *RestaurantService) isWithinOpeningHours(openingTime, closingTime string) bool {
	now := time.Now()
	currentTime := now.Format("15:04:05")

	if closingTime < openingTime {
		return currentTime >= openingTime || currentTime <= closingTime
	}

	return currentTime >= openingTime && currentTime <= closingTime
}
