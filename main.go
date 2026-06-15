package main

import (
	"github.com/VelVit24/projext/database"
	"github.com/VelVit24/projext/handlers"
	mw "github.com/VelVit24/projext/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.ConnDB()
	defer db.Close()
	h := handlers.Handler{DB: db}
	router.POST("/instruments", h.POSTInstruments)         // добавление инструмента
	router.PUT("/instruments/:id", h.PUTInstruments)       // изменение инструмента
	router.DELETE("/instruments/:id", h.DELETEInstruments) // удаление инструмента
	router.GET("/instruments", h.GETInstruments)           // получение всего инструмента

	router.POST("/register", h.POSTRegister) // авторизация пользователей
	router.POST("/login", h.POSTLogin)

	router.POST("/cart", mw.AuthMiddleware(), h.POSTCart)         // добавление в корзину инструмента id
	router.PUT("/cart", mw.AuthMiddleware(), h.PUTCart)           // изменение элемента корзины
	router.DELETE("/cart/:id", mw.AuthMiddleware(), h.DELETECart) // удаление элемента корзины
	router.GET("/cart", mw.AuthMiddleware(), h.GETCart)           // получение всей корзины

	router.POST("/orders", mw.AuthMiddleware()) // создание заказа пользователем
	router.PUT("/orders/:id")
	router.DELETE("/orders/:id")
	router.GET("/orders")

	router.POST("/categories", h.POSTCategories)         // добавление категории
	router.PUT("/categories/:id", h.PUTCategories)       // изменение категории
	router.DELETE("/categories/:id", h.DELETECategories) // удаление категории
	router.GET("/categories", h.GETCategories)           // получение всех категорий

	router.Run(":8080")
}
