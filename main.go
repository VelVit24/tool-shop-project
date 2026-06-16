package main

import (
	"github.com/VelVit24/projext/app"
	"github.com/VelVit24/projext/database"
	"github.com/VelVit24/projext/routes"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.ConnDB()
	defer db.Close()

	// m, err := migrate.New(
	// 	"file://database/migrations",
	// 	"postgres://postgres:080907@localhost:5432/tool-shop?sslmode=disable",
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := m.Up(); err != nil {
	// 	log.Fatal(err)
	// }

	app := app.New(db)
	routes.Setup(router, app)
	router.Run(":8080")
}
