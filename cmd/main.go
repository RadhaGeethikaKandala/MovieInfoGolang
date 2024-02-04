package main

import (
	"fmt"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/config"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/database"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/router"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables(".env")
}

func main() {

	// Get Server Configuration
	serverConf := config.ReadConfig().Server
	host := serverConf.Host
	port := serverConf.Port

	// Run DB migration
	database.RunDatabaseMigrations()

	// Run gin-engine/app
	engine := gin.Default()
	router.Router(engine)

	engine.Run(fmt.Sprintf("%s:%s", host, port))
}
