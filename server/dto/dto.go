package dto

import "github.com/VelVit24/projext/models"

type ProductsResponce struct {
	Page     int              `json:"page"`
	Limit    int              `json:"limit"`
	Products []models.Product `json:"products"`
	Total    int              `json:"total"`
}

type CartItems struct {
	Id_product int    `json:"id_product"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	Image_url  string `json:"image_url"`
	Amount     int    `json:"amount"`
	IsInStock  bool   `json:"is_in_stock"`
}
type OrderView struct {
	Order      models.Order `json:"order"`
	UserEmail  string       `json:"user_email"`
	OrderItems []CartItems  `json:"order_items"`
}

type ErrorResponce struct {
	Error string `json:"error"`
}
