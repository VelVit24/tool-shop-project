package app

import (
	"database/sql"

	"github.com/VelVit24/projext/handlers"
	"github.com/VelVit24/projext/repository"
	"github.com/VelVit24/projext/service"
)

type App struct {
	ProductHandler *handlers.ProductHandler
}

func New(db *sql.DB) *App {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	return &App{
		ProductHandler: productHandler,
	}
}
