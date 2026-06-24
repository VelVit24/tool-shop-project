package dto

import "github.com/VelVit24/projext/models"

type ProductsResponce struct {
	Page     int              `json:"page"`
	Limit    int              `json:"limit"`
	Products []models.Product `json:"products"`
	Total    int              `json:"total"`
}

type CartItems struct {
	Id_product  int    `json:"id_product"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Image_count int    `json:"image_count"`
	Amount      int    `json:"amount"`
	IsInStock   bool   `json:"is_in_stock"`
}

type ErrorResponce struct {
	Error string `json:"error"`
}

type OrderFull struct {
	Order      models.Order `json:"order"`
	User       models.User  `json:"user"`
	OrderItems []CartItems  `json:"cart_items"`
}
type OrderResponce struct {
	Orders []OrderFull `json:"orders"`
	Page   int         `json:"page"`
	Limit  int         `json:"limit"`
	Total  int         `json:"total"`
}

type OrderRequestNoAuth struct {
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	CartItems []CartItems `json:"cart_items"`
}
