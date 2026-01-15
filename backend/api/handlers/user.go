package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimashii-dan/food-delivery-app/api/domain"
	"github.com/kimashii-dan/food-delivery-app/services/user-service/pb"
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
		"refreshToken",
		grpcResp.RefreshToken,
		5*24*60*60,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, domain.LoginResponse{
		AccessToken: grpcResp.AccessToken,
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	c.SetCookie(
		"refreshToken",
		grpcResp.RefreshToken,
		5*24*60*60,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, domain.RefreshResponse{
		AccessToken: grpcResp.AccessToken,
	})
}
