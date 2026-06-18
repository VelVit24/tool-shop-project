package dto

type ProductFiler struct {
	Page       int
	Limit      int
	CategoryID *int
	PriceFrom  *int
	PriceTo    *int
	InStock    *bool
	Search     *string
	Sort       string
}
