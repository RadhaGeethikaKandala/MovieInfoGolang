package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnHelloWorldWhentheHandlerIsInvoked(t *testing.T) {

	engine := gin.Default()

	engine.GET("/test/hello", SayHello)

	request, err := http.NewRequest(http.MethodGet, "/test/hello", nil)
	require.NoError(t, err)

	responseRecoder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecoder, request)

	bytes, err := io.ReadAll(responseRecoder.Body)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecoder.Result().StatusCode)
	assert.Equal(t, "hello world!", string(bytes))

}

func TestShouldAllMoviesWhenGivenNameisSpiderMan(t *testing.T) {

	engine := gin.Default()

	engine.GET("/test/movieList/:name", GetMovieList)

	assertMovie(t, engine, "spiderman")

}

func TestShouldAllMoviesWhenGivenNameisBatMan(t *testing.T) {

	engine := gin.Default()

	engine.GET("/test/movieList/:name", GetMovieList)
	assertMovie(t, engine, "spiderman")

}

func assertMovie(t *testing.T, engine *gin.Engine, name string) {
	request, err := http.NewRequest(http.MethodGet, "/test/movieList/"+name, nil)
	require.NoError(t, err)

	responseRecoder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecoder, request)

	require.NoError(t, err)

	var ResponseMovieList []dto.Movie
	json.NewDecoder(responseRecoder.Body).Decode(&ResponseMovieList)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecoder.Result().StatusCode)

	for _, movie := range ResponseMovieList {
		assert.Contains(t, movie.Title, name)

	}
}
