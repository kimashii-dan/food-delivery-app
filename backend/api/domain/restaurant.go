package domain

type Restaurant struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Phone       string  `json:"phone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	LogoURL     string  `json:"logo_url"`
	OpeningTime string  `json:"opening_time"`
	ClosingTime string  `json:"closing_time"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type MenuItem struct {
	ID           string  `json:"id"`
	RestaurantID string  `json:"restaurant_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	ImageURL     string  `json:"image_url"`
	IsAvailable  bool    `json:"is_available"`
	Category     string  `json:"category"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GetRestaurantsResponse struct {
	Restaurants []*Restaurant `json:"restaurants"`
	Total       int32         `json:"total"`
}

type GetRestaurantResponse struct {
	Restaurant *Restaurant `json:"restaurant"`
}

type GetMenuResponse struct {
	Items []*MenuItem `json:"items"`
	Total int32       `json:"total"`
}

type GetMenuItemResponse struct {
	Item *MenuItem `json:"item"`
}

type GetRestaurantStatusResponse struct {
	IsAcceptingOrders bool   `json:"is_accepting_orders"`
	OpeningTime       string `json:"opening_time"`
	ClosingTime       string `json:"closing_time"`
}

type MenuItemValidation struct {
	ItemID      string `json:"item_id"`
	IsAvailable bool   `json:"is_available"`
	Name        string `json:"name"`
}

type ValidateMenuItemsRequest struct {
	RestaurantID string   `json:"restaurant_id" binding:"required"`
	ItemIDs      []string `json:"item_ids" binding:"required,min=1"`
}

type ValidateMenuItemsResponse struct {
	AllAvailable     bool                  `json:"all_available"`
	Items            []*MenuItemValidation `json:"items"`
	UnavailableItems []string              `json:"unavailable_items"`
}
