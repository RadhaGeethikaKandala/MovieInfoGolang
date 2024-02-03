package repository

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetMovies(t *testing.T) {

	t.Run("get all movies", func(t *testing.T) {
		db, mock := NewMock()

		repository := NewRepository(db)

		rows := sqlmock.NewRows([]string{"id", "title", "year", "rated", "released", "runtime", "genre", "director",
			"writer", "actors", "plot", "language", "country",
			"awards", "poster", "metascore", "imdbrating", "imdbvotes", "imdbid",
			"type", "dvd", "boxoffice", "production", "website", "response"}).
			AddRow("1", "Batman Begins", "2005", "PG-13", "15 Jun 2005", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "").
			AddRow("2", "Batman Rises", "2010", "PG-13", "16 Jun 2010", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "").
			AddRow("3", "Batman Returns", "2015", "PG-13", "16 Jun 2015", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM movies`)).WillReturnRows(rows)

		movies := repository.GetMovies(&request.MoviesRequest{})

		assert.Equal(t, 3, len(movies))
		assert.Equal(t, "Batman Begins", movies[0].Title)
		assert.Equal(t, "Batman Rises", movies[1].Title)
		assert.Equal(t, "Batman Returns", movies[2].Title)

		assert.Equal(t, "15 Jun 2005", movies[0].Released)
		assert.Equal(t, "16 Jun 2010", movies[1].Released)
		assert.Equal(t, "16 Jun 2015", movies[2].Released)

	})

	t.Run("get movies with 2 matching genres which is given", func(t *testing.T) {
		db, mock := NewMock()

		repository := NewRepository(db)

		rows := sqlmock.NewRows([]string{"id", "title", "year", "rated", "released", "runtime", "genre", "director",
			"writer", "actors", "plot", "language", "country",
			"awards", "poster", "metascore", "imdbrating", "imdbvotes", "imdbid",
			"type", "dvd", "boxoffice", "production", "website", "response"}).
			AddRow("1", "Batman Begins", "2005", "PG-13", "15 Jun 2005", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "").
			AddRow("2", "Batman Rises", "2010", "PG-13", "16 Jun 2010", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")

		queryString := "SELECT * FROM movies WHERE LOWER(genre) LIKE LOWER('%' || $1 || '%') AND LOWER(year)=$2"
		mock.ExpectQuery(regexp.QuoteMeta(queryString)).WillReturnRows(rows)

		movies := repository.GetMovies(&request.MoviesRequest{Genre: "action", Year: "2005"})

		assert.Equal(t, 2, len(movies))
		assert.Equal(t, "Batman Begins", movies[0].Title)
		assert.Equal(t, "Batman Rises", movies[1].Title)

		assert.Equal(t, "15 Jun 2005", movies[0].Released)
		assert.Equal(t, "16 Jun 2010", movies[1].Released)

	})

}

func TestGetRatingsForMovies(t *testing.T) {
	db, mock := NewMock()

	repository := NewRepository(db)

	rows := sqlmock.NewRows([]string{"id", "source", "value"}).
		AddRow("1", "Internet Movie Database", "8.2/10").
		AddRow("2", "Rotten Tomatoes", "85%").
		AddRow("3", "Metacritic", "70/100")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ratings`)).WillReturnRows(rows)

	ratings := repository.GetRatingsFor(1)

	assert.Equal(t, 3, len(ratings))
	assert.Equal(t, "Internet Movie Database", ratings[0].Source)
	assert.Equal(t, "Rotten Tomatoes", ratings[1].Source)
	assert.Equal(t, "Metacritic", ratings[2].Source)

	assert.Equal(t, "8.2/10", ratings[0].Value)
	assert.Equal(t, "85%", ratings[1].Value)
	assert.Equal(t, "70/100", ratings[2].Value)
}
