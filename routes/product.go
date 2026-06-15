package routes

import (
	"github.com/VelVit24/projext/app"
	"github.com/VelVit24/projext/handlers"
	mw "github.com/VelVit24/projext/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, app *app.App) {
	ProductRoutes(router, app.ProductHandler)
}

func ProductRoutes(router *gin.Engine, handler *handlers.ProductHandler) {
	admin := router.Group("/admin", mw.AuthMiddleware())
	{
		admin.POST("/products", handler.PostAdminProduct)             // добавление инструмента
		admin.PUT("/products/:id", handler.PutAdminProduct)           // изменение инструмента
		admin.DELETE("/products/:id", handler.DeleteAdminInstruments) // удаление инструмента
		admin.GET("/products", handler.GetAdminInstruments)           // получение всего инструмента
	}

}
