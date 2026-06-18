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
	router := gin.Default()
	router.Use(cors.Default())
	db := database.ConnDB()
	defer db.Close()
	app := app.New(db)
	routes.Setup(router, app)
	router.Run(":8080")
}
