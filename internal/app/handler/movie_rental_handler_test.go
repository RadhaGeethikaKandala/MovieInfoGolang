package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/response"
	mock_service "github.com/RadhaGeethikaKandala/MovieRental/internal/app/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnHelloWorldWhentheHandlerIsInvoked(t *testing.T) {

	engine := gin.Default()

	ctrl := gomock.NewController(t)
	service := mock_service.NewMockMovieService(ctrl)
	handler := NewHandler(service)

	engine.GET("/test/hello", handler.SayHello)

	request, err := http.NewRequest(http.MethodGet, "/test/hello", nil)
	require.NoError(t, err)

	responseRecoder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecoder, request)

	bytes, err := io.ReadAll(responseRecoder.Body)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecoder.Result().StatusCode)
	assert.Equal(t, "hello world!", string(bytes))

}

func TestGetMovies(t *testing.T) {

	engine := gin.Default()

	ctrl := gomock.NewController(t)
	service := mock_service.NewMockMovieService(ctrl)
	handler := NewHandler(service)
	// engine.GET("/api/movies/:name", handler.GetMovieList)
	engine.GET("/api/movies/", handler.GetMoviesFromDb)
	engine.GET("/api/movies/:imdbid", handler.GetMovieDetails)

	// t.Run("service should return movie lists when given name is spiderman", func(t *testing.T) {
	// 	movieList := []dto.Movie{
	// 		{
	// 			Title: "spiderman1",
	// 		},
	// 		{
	// 			Title: "spiderman2",
	// 		},
	// 	}

	// 	name := "spiderman"

	// 	service.EXPECT().GetMovies(name).Times(1).Return(movieList, nil)

	// 	responseStruct, code := MakeMoviesApiCall(name, t, engine, nil)

	// 	assertValidInput(t, responseStruct, name, "")
	// 	assert.Equal(t, http.StatusOK, code)

	// })

	// t.Run("service should return movie lists when given name is batman", func(t *testing.T) {
	// 	movieList := []dto.Movie{
	// 		{
	// 			Title: "batman",
	// 		},
	// 		{
	// 			Title: "batman returns",
	// 		},
	// 	}

	// 	name := "batman"

	// 	service.EXPECT().GetMovies(name).Times(1).Return(movieList, nil)

	// 	responseStruct, code := MakeMoviesApiCall(name, t, engine, nil)

	// 	assertValidInput(t, responseStruct, name, "")
	// 	assert.Equal(t, http.StatusOK, code)

	// })

	// t.Run("service should return error message when given name is not found", func(t *testing.T) {

	// 	name := "swdkenwenc"
	// 	errFromService := errors.New("Movies not found!")

	// 	service.EXPECT().GetMovies(name).Times(1).Return(nil, errFromService)

	// 	responseStruct, code := MakeMoviesApiCall(name, t, engine, nil)

	// 	assertValidInput(t, responseStruct, name, errFromService.Error())
	// 	assert.Equal(t, http.StatusOK, code)
	// })

	// t.Run("service should return error message when given name is empty", func(t *testing.T) {

	// 	name := " "
	// 	err := errors.New("name cannot be empty")

	// 	service.EXPECT().GetMovies(name).Times(0)

	// 	responseStruct, code := MakeMoviesApiCall(name, t, engine, nil)

	// 	assert.Equal(t, err.Error(), responseStruct.ErrorMessage)
	// 	assert.Equal(t, http.StatusBadRequest, code)

	// })

	t.Run("it should return movie lists from db", func(t *testing.T) {
		truncatedMovies := []response.TruncatedMovie{
			{
				Title: "spiderman1",
			},
			{
				Title: "spiderman2",
			},
		}

		movieList := response.TruncatedMovieReponse{Movies: truncatedMovies}
		service.EXPECT().GetMoviesFromDb(&request.MoviesRequest{}).Times(1).Return(movieList)
		responseStruct, code := MakeMoviesApiCall("", t, engine, nil)
		assertValidInput(t, responseStruct, "", "")
		assert.Equal(t, http.StatusOK, code)
	})

	t.Run("it should return matching movies with matching given genre, actor or year", func(t *testing.T) {
		truncatedMovies := []response.TruncatedMovie{
			{
				Title: "spiderman1",
				Genre: "Action",
			},
			{
				Title:  "spiderman2",
				Genre:  "Fantasy",
				Actors: "Robert",
			},
			{
				Title: "spiderman3",
				Genre: "Fantasy",
				Year:  "2007",
			},
		}
		movieList := response.TruncatedMovieReponse{Movies: truncatedMovies}
		mymap := make(map[string]interface{})
		mymap["genre"] = "Action"
		mymap["actors"] = "Robert"
		mymap["year"] = "2007"
		service.EXPECT().GetMoviesFromDb(&request.MoviesRequest{Genre: "Action", Actors: "Robert", Year: "2007"}).Times(1).Return(movieList)
		responseStruct, code := MakeMoviesApiCall("", t, engine, mymap)
		assertValidInput(t, responseStruct, "", "")
		assert.Equal(t, http.StatusOK, code)

	})

	t.Run("it should return movie details for correct imdbid", func(t *testing.T) {
		movie := dto.Movie{
			Title:  "spiderman2",
			Genre:  "Fantasy",
			Actors: "Robert",
		}
		movieResp := response.MovieResponse{Movie: movie}
		service.EXPECT().GetMovieDetails("imdb12345").Times(1).Return(movieResp, nil)
		responseStruct, code := MakeMoviesApiCall("imdb12345", t, engine, nil)
		assertValidInput(t, responseStruct, "", "")
		assert.Equal(t, http.StatusOK, code)

	})

	t.Run("it should return http status 400 if whitespaces are passed instead of imdbid for movie details call", func(t *testing.T) {
		responseStruct, code := MakeMoviesApiCall("    ", t, engine, nil)
		assertValidInput(t, responseStruct, "", "")
		assert.Equal(t, http.StatusBadRequest, code)

	})

	t.Run("it should return http status 400 if invalid imdbid for movie details call", func(t *testing.T) {
		movieResp := response.MovieResponse{}
		service.EXPECT().GetMovieDetails("imbd-invalid").Times(1).Return(movieResp, errors.New("Error"))
		responseStruct, code := MakeMoviesApiCall("imbd-invalid", t, engine, nil)
		assertValidInput(t, responseStruct, "", "")
		assert.Equal(t, http.StatusBadRequest, code)

	})
}

func TestCart(t *testing.T) {
	engine := gin.Default()

	ctrl := gomock.NewController(t)
	service := mock_service.NewMockMovieService(ctrl)
	handler := NewHandler(service)

	engine.POST("/api/movies/cart", handler.AddMovieToCart)

	t.Run("it should add movies to cart if user is valid and movie is valid", func(t *testing.T) {

		addToCartReq := request.AddToCartRequest{
			UserId: "1",
			ImdbId: "1",
		}

		service.EXPECT().AddMovieToCart(&addToCartReq).Times(1).Return(nil)

		marshalled, err := json.Marshal(addToCartReq)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/movies/cart", bytes.NewReader(marshalled))
		require.NoError(t, err)

		responseRecoder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecoder, request)

		var responseStruct response.ApiResponse
		json.NewDecoder(responseRecoder.Body).Decode(&responseStruct)
		require.NoError(t, err)

		val, _ := strconv.Atoi(responseStruct.Code)
		assert.Equal(t, http.StatusOK, val)
		assert.Equal(t, "Added movie to cart successfully", responseStruct.Message)
	})

	t.Run("it should add movies to cart if user or movie is invalid", func(t *testing.T) {

		addToCartReq := request.AddToCartRequest{
			UserId: "1invalid",
			ImdbId: "invalidimdb",
		}

		service.EXPECT().AddMovieToCart(&addToCartReq).Times(1).Return(errors.New("invalid user id"))

		marshalled, err := json.Marshal(addToCartReq)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/movies/cart", bytes.NewReader(marshalled))
		require.NoError(t, err)

		responseRecoder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecoder, request)

		var responseStruct response.ApiResponse
		json.NewDecoder(responseRecoder.Body).Decode(&responseStruct)
		require.NoError(t, err)

		val, _ := strconv.Atoi(responseStruct.Code)
		assert.Equal(t, http.StatusBadRequest, val)
		assert.Equal(t, "invalid user id", responseStruct.Message)
	})
}




func assertValidInput(t *testing.T, responseStruct response.TruncatedMovieReponse, name string, errMessage string) {

	for _, movie := range responseStruct.Movies {

		assert.Contains(t, movie.Title, name)

	}
	// } else {
	// 	assert.Equal(t, 0, len(responseStruct.MovieList))
	// 	assert.Equal(t, errMessage, responseStruct.ErrorMessage)
	// }
}

func MakeMoviesApiCall(pathParam string, t *testing.T, engine *gin.Engine, queryParams map[string]interface{}) (response.TruncatedMovieReponse, int) {
	url := CreateMoviesUrl(pathParam, queryParams)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	responseRecoder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecoder, request)

	var responseStruct response.TruncatedMovieReponse
	json.NewDecoder(responseRecoder.Body).Decode(&responseStruct)
	require.NoError(t, err)

	return responseStruct, responseRecoder.Result().StatusCode
}

func CreateMoviesUrl(pathParam string, queryParams map[string]interface{}) string {
	url := "/api/movies/" + pathParam
	if queryParams != nil {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = strings.TrimRight(url, "&")
	}
	return url
}
