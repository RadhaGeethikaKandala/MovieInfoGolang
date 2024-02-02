package router

import (
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/client"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/database"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/handler"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/repository"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/service"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {

	client := client.NewOmdbClient()
	repository := repository.NewRepository(database.CreateDatabaseConn())
	service := service.NewMovieService(client, repository)
	handler := handler.NewHandler(service)

	engine.GET("/hello", handler.SayHello)
	movieRentalApiGroup := engine.Group("/api/movies/")
	{
		movieRentalApiGroup.GET("/", handler.GetMoviesFromDb)
	}
}
