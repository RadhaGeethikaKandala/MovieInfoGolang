package handler

import (
	"strings"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/response"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/service"
	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	SayHello(ctx *gin.Context)
	GetMovieList(ctx *gin.Context)
	GetMoviesFromDb(ctx *gin.Context)
	GetMovieDetails(ctx *gin.Context)
}

type movieHandler struct {
	service service.MovieService
}

func NewHandler(service service.MovieService) MovieHandler {
	return &movieHandler{
		service: service,
	}
}

func (mh movieHandler) SayHello(ctx *gin.Context) {
	ctx.String(200, "hello world!")
}

func (mh movieHandler) GetMovieList(ctx *gin.Context) {

	movieName := ctx.Param("name")

	if strings.TrimSpace(movieName) == "" {
		ctx.JSON(400, gin.H{
			"movieList":    nil,
			"errorMessage": "name cannot be empty",
		})
		return
	}

	movieList, err := mh.service.GetMovies(movieName)

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}

	ctx.JSON(200, gin.H{
		"movieList":    movieList,
		"errorMessage": errorMessage,
	})

}

func (mh movieHandler) GetMoviesFromDb(ctx *gin.Context) {
	var moviesRequest request.MoviesRequest
	ctx.ShouldBindQuery(&moviesRequest)

	movies := mh.service.GetMoviesFromDb(&moviesRequest)
	ctx.JSON(200, gin.H{
		"movies": movies,
	})
}

func (mh movieHandler) GetMovieDetails(ctx *gin.Context) {
	// get imdbid
	imdbid := ctx.Param("imdbid")
	if strings.TrimSpace(imdbid) == "" {
		ctx.JSON(400, response.ErrorResponse{
			Status:  "error",
			Message: "name cannot be empty or have whitespaces",
			Code:    "400",
		})
		return
	}

	movie := mh.service.GetMovieDetails(imdbid)
	ctx.JSON(200, gin.H{
		"movie_detail": movie,
	})

}
