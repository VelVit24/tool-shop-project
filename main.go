package main

import (
	"github.com/VelVit24/projext/app"
	"github.com/VelVit24/projext/database"
	"github.com/VelVit24/projext/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.ConnDB()
	defer db.Close()
	app := app.New(db)
	routes.Setup(router, app)

	router.POST("/register", h.PostRegister) // авторизация пользователей
	router.POST("/login", h.PostLogin)

	router.POST("/cart", mw.AuthMiddleware(), h.PostCart)           // добавление в корзину инструмента id
	router.PUT("/cart", mw.AuthMiddleware(), h.PutCart)             // изменение элемента корзины
	router.DELETE("/cart/:id", mw.AuthMiddleware(), h.DeleteCartId) // удаление элемента корзины
	router.GET("/cart", mw.AuthMiddleware(), h.GetCart)             // получение всей корзины

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
