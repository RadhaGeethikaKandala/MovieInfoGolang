package main

import (

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/config"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/database"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/router"
	"github.com/gin-gonic/gin"

)

func init() {
	config.LoadEnvVariables(".env")
}

func main() {

	// Run DB migration
	database.RunDatabaseMigrations()

	// Run gin-engine/app
	engine := gin.Default()
	router.Router(engine)

	engine.Run(":8085")
}
