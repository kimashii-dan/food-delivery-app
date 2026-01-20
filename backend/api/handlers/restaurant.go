package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kimashii-dan/food-delivery-app/backend/api/domain"
	"github.com/kimashii-dan/food-delivery-app/backend/services/restaurant-service/pb"
)

type RestaurantHandler struct {
	restaurantClient pb.RestaurantServiceClient
}

func NewRestaurantHandler(restaurantClient pb.RestaurantServiceClient) *RestaurantHandler {
	return &RestaurantHandler{
		restaurantClient: restaurantClient,
	}
}

func (h *RestaurantHandler) GetRestaurants(c *gin.Context) {
	page := int32(1)
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = int32(p)
		}
	}

	grpcReq := &pb.GetRestaurantsRequest{
		Page: page,
	}

	grpcResp, err := h.restaurantClient.GetRestaurants(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	restaurants := make([]*domain.Restaurant, len(grpcResp.Restaurants))
	for i, r := range grpcResp.Restaurants {
		restaurants[i] = &domain.Restaurant{
			ID:          r.Id,
			Name:        r.Name,
			Description: r.Description,
			Address:     r.Address,
			Phone:       r.Phone,
			Latitude:    r.Latitude,
			Longitude:   r.Longitude,
			LogoURL:     r.LogoUrl,
			OpeningTime: r.OpeningTime,
			ClosingTime: r.ClosingTime,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, domain.GetRestaurantsResponse{
		Restaurants: restaurants,
		Total:       grpcResp.Total,
	})
}

func (h *RestaurantHandler) GetRestaurant(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "restaurant id is required"})
		return
	}

	grpcReq := &pb.GetRestaurantRequest{
		Id: id,
	}

	grpcResp, err := h.restaurantClient.GetRestaurant(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	r := grpcResp.Restaurant
	c.JSON(http.StatusOK, domain.GetRestaurantResponse{
		Restaurant: &domain.Restaurant{
			ID:          r.Id,
			Name:        r.Name,
			Description: r.Description,
			Address:     r.Address,
			Phone:       r.Phone,
			Latitude:    r.Latitude,
			Longitude:   r.Longitude,
			LogoURL:     r.LogoUrl,
			OpeningTime: r.OpeningTime,
			ClosingTime: r.ClosingTime,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		},
	})
}

func (h *RestaurantHandler) GetMenu(c *gin.Context) {
	restaurantID := c.Param("id")
	if restaurantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "restaurant id is required"})
		return
	}

	page := int32(1)
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = int32(p)
		}
	}

	grpcReq := &pb.GetMenuRequest{
		RestaurantId: restaurantID,
		Page:         page,
	}

	grpcResp, err := h.restaurantClient.GetMenu(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items := make([]*domain.MenuItem, len(grpcResp.Items))
	for i, item := range grpcResp.Items {
		items[i] = &domain.MenuItem{
			ID:           item.Id,
			RestaurantID: item.RestaurantId,
			Name:         item.Name,
			Description:  item.Description,
			Price:        item.Price,
			ImageURL:     item.ImageUrl,
			IsAvailable:  item.IsAvailable,
			Category:     item.Category,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, domain.GetMenuResponse{
		Items: items,
		Total: grpcResp.Total,
	})
}

func (h *RestaurantHandler) GetMenuItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "menu item id is required"})
		return
	}

	grpcReq := &pb.GetMenuItemRequest{
		Id: id,
	}

	grpcResp, err := h.restaurantClient.GetMenuItem(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "menu item not found"})
		return
	}

	item := grpcResp.Item
	c.JSON(http.StatusOK, domain.GetMenuItemResponse{
		Item: &domain.MenuItem{
			ID:           item.Id,
			RestaurantID: item.RestaurantId,
			Name:         item.Name,
			Description:  item.Description,
			Price:        item.Price,
			ImageURL:     item.ImageUrl,
			IsAvailable:  item.IsAvailable,
			Category:     item.Category,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		},
	})
}

func (h *RestaurantHandler) GetRestaurantStatus(c *gin.Context) {
	restaurantID := c.Param("id")
	if restaurantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "restaurant id is required"})
		return
	}

	grpcReq := &pb.GetRestaurantStatusRequest{
		RestaurantId: restaurantID,
	}

	grpcResp, err := h.restaurantClient.GetRestaurantStatus(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, domain.GetRestaurantStatusResponse{
		IsAcceptingOrders: grpcResp.IsAcceptingOrders,
		OpeningTime:       grpcResp.OpeningTime,
		ClosingTime:       grpcResp.ClosingTime,
	})
}

func (h *RestaurantHandler) ValidateMenuItems(c *gin.Context) {
	var req domain.ValidateMenuItemsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.ValidateMenuItemsRequest{
		RestaurantId: req.RestaurantID,
		ItemIds:      req.ItemIDs,
	}

	grpcResp, err := h.restaurantClient.ValidateMenuItems(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items := make([]*domain.MenuItemValidation, len(grpcResp.Items))
	for i, item := range grpcResp.Items {
		items[i] = &domain.MenuItemValidation{
			ItemID:      item.ItemId,
			IsAvailable: item.IsAvailable,
			Name:        item.Name,
		}
	}

	c.JSON(http.StatusOK, domain.ValidateMenuItemsResponse{
		AllAvailable:     grpcResp.AllAvailable,
		Items:            items,
		UnavailableItems: grpcResp.UnavailableItems,
	})
}
