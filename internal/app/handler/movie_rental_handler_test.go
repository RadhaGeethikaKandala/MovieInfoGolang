package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
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
	engine.GET("/api/movies/", handler.GetMoviesFromDb)
	engine.GET("/api/movies/:name", handler.GetMovieList)

	t.Run("service should return movie lists when given name is spiderman", func(t *testing.T) {
		movieList := []dto.Movie{
			{
				Title: "spiderman1",
			},
			{
				Title: "spiderman2",
			},
		}

		name := "spiderman"

		service.EXPECT().GetMovies(name).Times(1).Return(movieList, nil)

		responseStruct, code := MakeApiCall(name, t, engine, nil)

		assertValidInput(t, responseStruct, name, "")
		assert.Equal(t, http.StatusOK, code)

	})

	t.Run("service should return movie lists when given name is batman", func(t *testing.T) {
		movieList := []dto.Movie{
			{
				Title: "batman",
			},
			{
				Title: "batman returns",
			},
		}

		name := "batman"

		service.EXPECT().GetMovies(name).Times(1).Return(movieList, nil)

		responseStruct, code := MakeApiCall(name, t, engine, nil)

		assertValidInput(t, responseStruct, name, "")
		assert.Equal(t, http.StatusOK, code)

	})

	t.Run("service should return error message when given name is not found", func(t *testing.T) {

		name := "swdkenwenc"
		errFromService := errors.New("Movies not found!")

		service.EXPECT().GetMovies(name).Times(1).Return(nil, errFromService)

		responseStruct, code := MakeApiCall(name, t, engine, nil)

		assertValidInput(t, responseStruct, name, errFromService.Error())
		assert.Equal(t, http.StatusOK, code)
	})

	t.Run("service should return error message when given name is empty", func(t *testing.T) {

		name := " "
		err := errors.New("name cannot be empty")

		service.EXPECT().GetMovies(name).Times(0)

		responseStruct, code := MakeApiCall(name, t, engine, nil)

		assert.Equal(t, err.Error(), responseStruct.ErrorMessage)
		assert.Equal(t, http.StatusBadRequest, code)

	})

	t.Run("service should return movie lists from db", func(t *testing.T) {
		movieList := []dto.Movie{
			{
				Title: "spiderman1",
			},
			{
				Title: "spiderman2",
			},
		}


		service.EXPECT().GetMoviesFromDb(&request.MoviesRequest{}).Times(1).Return(movieList)
		responseStruct, code := MakeApiCall("", t, engine, nil)
		assertValidInput(t, responseStruct, "", "")
		assert.Equal(t, http.StatusOK, code)
	})

	t.Run("service should return matching movies with matching given genre, actor or year", func(t *testing.T) {
		movieList := []dto.Movie{
			{
				Title: "spiderman1",
				Genre: "Action",

			},
			{
				Title: "spiderman2",
				Genre: "Fantasy",
				Actors: "Robert",
			},
			{
				Title: "spiderman3",
				Genre: "Fantasy",
				Year: "2007",
			},
		}
		service.EXPECT().GetMoviesFromDb(&request.MoviesRequest{Genre:"Action", Actors: "Robert", Year: "2007"}).Times(1).Return(movieList)
		responseStruct, code := MakeApiCall("", t, engine, map[string]string{"genre":"Action", "actors":"Robert", "year": "2007"})
		assertValidInput(t, responseStruct, "", "")
		assert.Equal(t, http.StatusOK, code)

	})
}

func assertValidInput(t *testing.T, responseStruct dto.MovieRentalResponse, name string, errMessage string) {

	if len(responseStruct.ErrorMessage) == 0 {
		for _, movie := range responseStruct.MovieList {

			assert.Contains(t, movie.Title, name)

		}
	} else {
		assert.Equal(t, 0, len(responseStruct.MovieList))
		assert.Equal(t, errMessage, responseStruct.ErrorMessage)
	}
}

func MakeApiCall(name string, t *testing.T, engine *gin.Engine, queryParams map[string]string) (dto.MovieRentalResponse, int) {
	url := CreateUrl(name, queryParams)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	responseRecoder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecoder, request)

	var responseStruct dto.MovieRentalResponse
	json.NewDecoder(responseRecoder.Body).Decode(&responseStruct)
	require.NoError(t, err)

	return responseStruct, responseRecoder.Result().StatusCode
}


func CreateUrl(name string, queryParams map[string]string) string{
	url := "/api/movies/"+name
	if queryParams != nil {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
	}
	return url
}