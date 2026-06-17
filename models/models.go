package models

import "time"

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Image_url   string `json:"image_url"`
	Id_category int    `json:"id_category"`
}
type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Cart struct {
	Id_product int `json:"id_product"`
	Amount     int `json:"amount"`
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
type Order struct {
	Id        int       `json:"id"`
	Status    string    `json:"status"`
	Total     int       `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}
type OrderItems struct {
	Id_product int `json:"id_product"`
	Amount     int `json:"amount"`
	Price      int `json:"price"`
}
type OrderView struct {
	Order      Order       `json:"order"`
	UserEmail  string      `json:"user_email"`
	OrderItems []CartItems `json:"order_items"`
}
