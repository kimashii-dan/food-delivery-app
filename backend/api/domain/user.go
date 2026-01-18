package domain

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=customer restaurant courier"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	User User `json:"user"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type Address struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	PostalCode string  `json:"postal_code"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	IsDefault  bool    `json:"is_default"`
	CreatedAt  string  `json:"created_at"`
}

type AddAddressRequest struct {
	Street     string  `json:"street" binding:"required"`
	City       string  `json:"city" binding:"required"`
	PostalCode string  `json:"postal_code" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
	IsDefault  bool    `json:"is_default" binding:"required"`
}

type AddAddressResponse struct {
	AddressId string `json:"address_id"`
}

type GetAddressesResponse struct {
	Addresses []*Address `json:"addresses"`
}
