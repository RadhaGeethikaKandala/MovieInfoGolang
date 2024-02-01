package router

import (
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/client"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/handler"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/service"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {

	client := client.NewOmdbClient()
	service := service.NewMovieService(client)
	handler := handler.NewHandler(service)

	engine.GET("/hello", handler.SayHello)
	omdbMovieApiGroup := engine.Group("/api/movies/")
	{
		omdbMovieApiGroup.GET(":name", handler.GetMovieList)
	}
}
