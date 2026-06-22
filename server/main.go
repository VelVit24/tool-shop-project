package main

import (
	"github.com/VelVit24/projext/app"
	"github.com/VelVit24/projext/database"
	"github.com/VelVit24/projext/routes"
	"github.com/gin-contrib/cors"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/gin-gonic/gin"
)

func main() {
	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	}
	router := gin.Default()
	router.Use(cors.New(config))
	db := database.ConnDB()
	defer db.Close()
	router.Static("/static", "./static")
	app := app.New(db)
	routes.Setup(router, app)
	router.Run(":8080")
}
