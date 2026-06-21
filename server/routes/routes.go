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
	OrderRoutes(router, app.OrderHandler)
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
	r.GET("/check/email", h.GetCheckEmail) // проверка уникальности email
	r.GET("/check/phone", h.GetCheckPhone) // проверка уникальности phone
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
	r.POST("/cart", mw.AuthMiddleware(), h.PostCart)         // добавление в корзину инструмента id
	r.PUT("/cart", mw.AuthMiddleware(), h.PutCart)           // изменение элемента корзины
	r.DELETE("/cart/:id", mw.AuthMiddleware(), h.DeleteCart) // удаление элемента корзины
	r.GET("/cart", mw.AuthMiddleware(), h.GetCart)           // получение всей корзины
}

func OrderRoutes(r *gin.Engine, h *handlers.OrderHandler) {
	r.POST("/orders", mw.AuthMiddleware(), h.PostOrders) // создание заказа пользователем
	r.GET("/orders", mw.AuthMiddleware(), h.GetOrders)
	admin := r.Group("/admin", mw.AuthMiddleware(), mw.CheckAdminMiddleware())
	{
		admin.DELETE("/orders/:id", h.DeleteAdminOrders)
		admin.PUT("/orders/:id", h.PutAdminOrders)
		admin.GET("/orders", mw.AuthMiddleware(), mw.CheckAdminMiddleware(), h.GetAdminOrders)
	}
}
