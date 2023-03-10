package productmodel

type NewProductInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Slug        *string  `json:"slug"`
	Code        *string  `json:"code"`
	Image       *string  `json:"image"`
	Images      *string  `json:"images"`
	Price       float64  `json:"price"`
	PriceMin    *float64 `json:"price_min"`
	PriceMax    *float64 `json:"price_max"`
}
