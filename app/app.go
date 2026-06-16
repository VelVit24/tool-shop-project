package app

import (
	"database/sql"

	"github.com/VelVit24/projext/handlers"
	"github.com/VelVit24/projext/repository"
	"github.com/VelVit24/projext/service"
)

type App struct {
	ProductHandler  *handlers.ProductHandler
	CategoryHandler *handlers.CategoryHandler
	AuthHandler     *handlers.AuthHandler
	CartHandler     *handlers.CartHandler
}

func New(db *sql.DB) *App {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	return &App{
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
		CartHandler:     cartHandler,
		AuthHandler:     authHandler,
	}
}
