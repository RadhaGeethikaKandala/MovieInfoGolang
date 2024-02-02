package main

import (
	"log"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/config"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/router"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

)

func init() {
	config.LoadEnvVariables(".env")
}

func main() {

	// Run DB migration

	m, err := migrate.New("file://internal/app/database/migrations", "postgres://movierental:movierental@localhost:5432/movie_rental?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	// Run gin-engine/app

	engine := gin.Default()
	router.Router(engine)

	engine.Run(":8085")
}
