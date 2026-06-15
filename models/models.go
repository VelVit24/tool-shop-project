package models

type Instrument struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Id_category int    `json:"id_cat"`
}
type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Cart struct {
	Id_instr int `json:"id_instr"`
	Amount   int `json:"amount"`
}
type Order struct {
	Id   int    `json:"id"`
	Desc string `json:"desc"`
}
