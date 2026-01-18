package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimashii-dan/food-delivery-app/backend/api/domain"
	"github.com/kimashii-dan/food-delivery-app/backend/services/user-service/pb"
)

type UserHandler struct {
	userClient pb.UserServiceClient
}

func NewUserHandler(userClient pb.UserServiceClient) *UserHandler {
	return &UserHandler{
		userClient: userClient,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Phone:    req.Phone,
		Role:     req.Role,
	}

	grpcResp, err := h.userClient.Register(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.RegisterResponse{
		UserID: grpcResp.UserId,
	})
}

func (h *UserHandler) Login(c *gin.Context) {

	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	grpcResp, err := h.userClient.Login(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.SetCookie(
		"accessToken",
		grpcResp.AccessToken,
		15*60,
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refreshToken",
		grpcResp.RefreshToken,
		5*24*60*60,
		"/api/users/refresh",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, domain.LoginResponse{
		User: domain.User{
			Id:        grpcResp.User.Id,
			Email:     grpcResp.User.Email,
			Name:      grpcResp.User.Name,
			Phone:     grpcResp.User.Phone,
			Role:      grpcResp.User.Role,
			CreatedAt: grpcResp.User.CreatedAt,
		},
	})
}

func (h *UserHandler) Refresh(c *gin.Context) {

	tokenString, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	grpcReq := &pb.RefreshRequest{
		RefreshToken: tokenString,
	}
	grpcResp, err := h.userClient.Refresh(c.Request.Context(), grpcReq)
	if err != nil {
		c.SetCookie("accessToken", "", -1, "/", "", false, true)
		c.SetCookie("refreshToken", "", -1, "/api/users/refresh", "", false, true)

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	c.SetCookie(
		"accessToken",
		grpcResp.AccessToken,
		15*60,
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refreshToken",
		grpcResp.RefreshToken,
		5*24*60*60,
		"/api/users/refresh",
		"",
		false,
		true,
	)

	c.Status(http.StatusOK)
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "", false, true)
	c.SetCookie("refreshToken", "", -1, "/api/users/refresh", "", false, true)

	c.Status(http.StatusOK)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	grpcReq := &pb.GetUserRequest{
		UserId: userID,
	}

	grpcResp, err := h.userClient.GetUser(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	c.JSON(http.StatusOK, domain.GetUserResponse{
		User: &domain.User{
			Id:        grpcResp.User.Id,
			Email:     grpcResp.User.Email,
			Name:      grpcResp.User.Name,
			Phone:     grpcResp.User.Phone,
			Role:      grpcResp.User.Role,
			CreatedAt: grpcResp.User.CreatedAt,
		},
	})
}

func (h *UserHandler) AddAddress(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	var req domain.AddAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.AddAddressRequest{
		UserId:     userID,
		Street:     req.Street,
		City:       req.City,
		PostalCode: req.PostalCode,
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		IsDefault:  req.IsDefault,
	}

	grpcResp, err := h.userClient.AddAddress(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.AddAddressResponse{
		AddressId: grpcResp.AddressId,
	})
}

func (h *UserHandler) GetAddresses(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	grpcReq := &pb.GetAddressesRequest{
		UserId: userID,
	}

	grpcResp, err := h.userClient.GetAddresses(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	addresses := make([]*domain.Address, len(grpcResp.Addresses))
	for i, addr := range grpcResp.Addresses {
		addresses[i] = &domain.Address{
			ID:         addr.Id,
			UserID:     addr.UserId,
			Street:     addr.Street,
			City:       addr.City,
			PostalCode: addr.PostalCode,
			Latitude:   addr.Latitude,
			Longitude:  addr.Longitude,
			IsDefault:  addr.IsDefault,
			CreatedAt:  addr.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, domain.GetAddressesResponse{
		Addresses: addresses,
	})
}
