package models

import "time"

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Image_url   string  `json:"image_url"`
	Id_category int     `json:"id_category"`
}
type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}
type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type Cart struct {
	Id_product int `json:"id_product"`
	Amount     int `json:"amount"`
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
