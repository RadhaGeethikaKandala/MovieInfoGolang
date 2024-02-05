package repository

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestGetMovie(t *testing.T) {
	db, mock := NewMock()
	repository := NewRepository(db)

	t.Run("it should return movie when imdb is valid", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "title", "year", "rated", "released", "runtime", "genre", "director",
			"writer", "actors", "plot", "language", "country",
			"awards", "poster", "metascore", "imdbrating", "imdbvotes", "imdbid",
			"type", "dvd", "boxoffice", "production", "website", "response"}).
			AddRow("1", "Batman Begins", "2005", "PG-13", "15 Jun 2005", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")

		queryString := "SELECT * FROM movies WHERE imdbid=$1"
		mock.ExpectQuery(regexp.QuoteMeta(queryString)).WillReturnRows(row)

		movie := repository.GetMovie("1234")

		assert.Equal(t, "Batman Begins", movie.Title)
		assert.Equal(t, "2005", movie.Year)
		assert.Equal(t, "PG-13", movie.Rated)
		assert.Equal(t, "15 Jun 2005", movie.Released)
	})

	t.Run("it should return empty movie when imdb is invalid", func(t *testing.T) {
		queryString := "SELECT * FROM movies WHERE imdbid=$1"
		mock.ExpectQuery(regexp.QuoteMeta(queryString)).WillReturnError(sql.ErrNoRows)

		movie := repository.GetMovie("1234")

		assert.Equal(t, "", movie.Title)
		assert.Equal(t, "", movie.Year)
		assert.Equal(t, "", movie.Rated)
		assert.Equal(t, "", movie.Released)
		assert.Equal(t, 0, movie.Id)
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

func TestAddMovieToCart(t *testing.T) {

	db, mock := NewMock()

	repository := NewRepository(db)

	t.Run("should return no of rows affected and no error", func(t *testing.T) {

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO cart`)).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repository.AddMovieToCart(1, 1)

		require.NoError(t, err)
	})

	t.Run("should return error when DB returns error", func(t *testing.T) {

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO cart`)).WithArgs(1, 12).WillReturnError(errors.New("Error while inserting into DB"))

		err := repository.AddMovieToCart(1, 12)

		assert.EqualError(t, err, "Error while inserting into DB")
	})

	t.Run("should return movie already added to cart error when DB returns unique primary key violation error", func(t *testing.T) {

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO cart`)).WithArgs(1, 12).WillReturnError(errors.New(`duplicate key value violates unique constraint "pk_cart"`))

		err := repository.AddMovieToCart(1, 12)

		assert.EqualError(t, err, "movie already added to cart")
	})

}

func TestGetCustomer(t *testing.T) {
	db, mock := NewMock()

	repository := NewRepository(db)

	t.Run("it should return customers when id is valid", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow("1", "arkaprabha", "arkaprabha@gmail.com")

		queryString := "SELECT * FROM customers WHERE id=$1"
		mock.ExpectQuery(regexp.QuoteMeta(queryString)).WillReturnRows(row)

		customer := repository.GetCustomer("1")

		assert.Equal(t, "arkaprabha", customer.Name)
		assert.Equal(t, "arkaprabha@gmail.com", customer.Email)
	})

	t.Run("it should empty cutsomer when id is not found in db", func(t *testing.T) {

		queryString := "SELECT * FROM customers WHERE id=$1"
		mock.ExpectQuery(regexp.QuoteMeta(queryString)).WillReturnError(sql.ErrNoRows)

		customer := repository.GetCustomer("1")

		assert.Equal(t, "", customer.Name)
		assert.Equal(t, "", customer.Email)
		assert.Equal(t, 0, customer.Id)
	})

}
