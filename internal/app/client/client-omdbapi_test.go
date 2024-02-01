package client

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/config"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.LoadEnvVariables("../../../.env")

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestShouldReturnListOfMoviesWhenGivenAName(t *testing.T) {

	responseStatus, movieList, responseError := GetOmdbResponse("spiderman")

	assert.True(t, responseStatus)
	assert.Equal(t, 0, len(responseError))

	assert.Greater(t, len(movieList), 0)
	for _, movie := range movieList {
		assert.Contains(t, strings.ToLower(movie.Title), "spiderman")

	}
}

func TestShouldErrorResponseWhenGivenNoNameIsGiven(t *testing.T) {

	responseStatus, movieList, responseError := GetOmdbResponse("")

	assert.False(t, responseStatus)
	assert.Equal(t, "Incorrect IMDb ID.", responseError)

	assert.Equal(t, len(movieList), 0)

}

func GetOmdbResponse(name string) (bool, []dto.Movie, string) {
	client := NewOmdbClient()
	omdbResponse := client.GetMovieList(name)
	responseStatus, _ := strconv.ParseBool(omdbResponse.Response)
	movieList := omdbResponse.Search
	return responseStatus, movieList, omdbResponse.Error
}