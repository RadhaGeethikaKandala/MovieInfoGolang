package handler

import (
	"fmt"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {

	ctx.String(200, "hello world!")
}

func GetMovieList(ctx *gin.Context) {
	fmt.Println(ctx.Param("name"))
	var movieList []dto.Movie
	movieList = append(movieList, dto.Movie{
		Title: ctx.Param("name"),
	}, dto.Movie{
		Title: ctx.Param("name") + " returns",
	})
	ctx.JSON(200, movieList)

}
