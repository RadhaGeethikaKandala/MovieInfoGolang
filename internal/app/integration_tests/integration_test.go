package testing

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/response"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/handler"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/repository"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDbInstance *sql.DB

func TestMain(m *testing.M) {
	testDB := SetupTestDatabase()
	testDbInstance = testDB.DbInstance
	defer testDB.TearDown()
	os.Exit(m.Run())
}

func TestGetAllMovies(t *testing.T) {
	repo := repository.NewRepository(testDbInstance)
	service := service.NewMovieService(repo)
	h := handler.NewHandler(service)

	engine := gin.Default()

	engine.GET("/test/movies", h.GetMoviesFromDb)

	request, err := http.NewRequest(http.MethodGet, "/test/movies", nil)
	require.NoError(t, err)

	responseRecoder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecoder, request)

	var responseStruct response.TruncatedMovieReponse
	json.NewDecoder(responseRecoder.Body).Decode(&responseStruct)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecoder.Result().StatusCode)
	for _, movie := range responseStruct.Movies {

		assert.Contains(t, movie.Title, "Batman")

		//check response
	}

}
