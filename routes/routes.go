package routes

import (
	"github.com/VelVit24/projext/app"
	"github.com/VelVit24/projext/handlers"
	mw "github.com/VelVit24/projext/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, app *app.App) {
	ProductRoutes(router, app.ProductHandler)
	AuthRoutes(router, app.AuthHandler)
	CategoryRoutes(router, app.CategoryHandler)
	CartRoutes(router, app.CartHandler)
}

func ProductRoutes(r *gin.Engine, h *handlers.ProductHandler) {
	admin := r.Group("/admin", mw.AuthMiddleware())
	{
		admin.POST("/products", h.PostAdminProduct)          // добавление инструмента
		admin.PUT("/products/:id", h.PutAdminProduct)        // изменение инструмента
		admin.DELETE("/products/:id", h.DeleteAdminProducts) // удаление инструмента
	}
	r.GET("/products/:id", h.GetProductsId) // получение инструмента по id
	r.GET("/products", h.GetProducts)       // получение всего инструмента
}

func AuthRoutes(r *gin.Engine, h *handlers.AuthHandler) {
	r.POST("/register", h.PostRegister) // авторизация пользователей
	r.POST("/login", h.PostLogin)
}

func CategoryRoutes(r *gin.Engine, h *handlers.CategoryHandler) {
	admin := r.Group("/admin", mw.AuthMiddleware(), mw.CheckAdminMiddleware())
	{
		admin.POST("/categories", h.PostAdminCategories)         // добавление категории
		admin.PUT("/categories/:id", h.PutAdminCategories)       // изменение категории
		admin.DELETE("/categories/:id", h.DeleteAdminCategories) // удаление категории
	}
	r.GET("/categories", h.GetCategories)
}

func CartRoutes(r *gin.Engine, h *handlers.CartHandler) {
	client := r.Group("", mw.AuthMiddleware())
	{
		client.POST("/cart", h.PostCart)         // добавление в корзину инструмента id
		client.PUT("/cart", h.PutCart)           // изменение элемента корзины
		client.DELETE("/cart/:id", h.DeleteCart) // удаление элемента корзины
		client.GET("/cart", h.GetCart)           // получение всей корзины
	}
}

/*

	router.POST("/orders", mw.AuthMiddleware()) // создание заказа пользователем
	router.PUT("/orders/:id")
	router.DELETE("/orders/:id")
	router.GET("/orders")

*/
