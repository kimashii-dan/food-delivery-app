package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kimashii-dan/food-delivery-app/services/user-service/pb"
	"github.com/kimashii-dan/food-delivery-app/services/user-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userRepo    *repository.UserRepository
	addressRepo *repository.AddressRepository
	jwtService  *JWTService
}

func NewUserService(userRepo *repository.UserRepository, addressRepo *repository.AddressRepository, jwtService *JWTService) *UserService {
	return &UserService{
		userRepo:    userRepo,
		addressRepo: addressRepo,
		jwtService:  jwtService,
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	exists, err := s.userRepo.EmailExists(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &repository.User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Name:         req.Name,
		Phone:        req.Phone,
		Role:         req.Role,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &pb.RegisterResponse{
		UserId: user.ID,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	accessDuration := 15 * time.Minute
	accessToken, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Role, accessDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	refreshDuration := 5 * 24 * time.Hour
	refreshToken, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Role, refreshDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: &pb.User{
			Id:    user.ID,
			Email: user.Email,
			Name:  user.Name,
			Phone: user.Phone,
			Role:  user.Role,
		},
	}, nil
}

func (s *UserService) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	tokenString := req.RefreshToken
	claims, err := s.jwtService.ValidateToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}

	accessDuration := 15 * time.Minute
	accessToken, err := s.jwtService.GenerateToken(claims.UserID, claims.Email, claims.Role, accessDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	refreshDuration := 5 * 24 * time.Hour
	refreshToken, err := s.jwtService.GenerateToken(claims.UserID, claims.Email, claims.Role, refreshDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Email: user.Email,
			Name:  user.Name,
			Phone: user.Phone,
			Role:  user.Role,
		},
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if err := s.userRepo.Update(ctx, req.UserId, req.Name, req.Phone); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	user, err := s.userRepo.GetByID(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated user: %w", err)
	}

	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Email: user.Email,
			Name:  user.Name,
			Phone: user.Phone,
			Role:  user.Role,
		},
	}, nil
}

func (s *UserService) AddAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	address := &repository.Address{
		ID:         uuid.New().String(),
		UserID:     req.UserId,
		Street:     req.Street,
		City:       req.City,
		PostalCode: req.PostalCode,
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		IsDefault:  req.IsDefault,
	}

	if err := s.addressRepo.Create(ctx, address); err != nil {
		return nil, fmt.Errorf("failed to create address: %w", err)
	}

	return &pb.AddAddressResponse{
		AddressId: address.ID,
	}, nil
}

func (s *UserService) GetAddresses(ctx context.Context, req *pb.GetAddressesRequest) (*pb.GetAddressesResponse, error) {
	addresses, err := s.addressRepo.GetByUserID(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get addresses: %w", err)
	}

	pbAddresses := make([]*pb.Address, len(addresses))
	for i, addr := range addresses {
		pbAddresses[i] = &pb.Address{
			Id:         addr.ID,
			UserId:     addr.UserID,
			Street:     addr.Street,
			City:       addr.City,
			PostalCode: addr.PostalCode,
			Latitude:   addr.Latitude,
			Longitude:  addr.Longitude,
			IsDefault:  addr.IsDefault,
		}
	}

	return &pb.GetAddressesResponse{
		Addresses: pbAddresses,
	}, nil
}
