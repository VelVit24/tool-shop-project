package dto

type ProductFiler struct {
	Page         int
	Limit        int
	CategorySlug *string
	PriceFrom    *int
	PriceTo      *int
	InStock      *bool
	Search       *string
	Sort         string
}
