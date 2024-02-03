package repository

import (
	"testing"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuery(t *testing.T) {
	t.Run("create get all movies query", func(t *testing.T) {
		request := &request.MoviesRequest{}
		query, params := buildMovieRentalQuery(request)

		assert.Equal(t, "SELECT * FROM movies", query)
		assert.Equal(t, 0, len(params))
	})

	t.Run("create get all movies query with matching given genre", func(t *testing.T) {

		request := &request.MoviesRequest{ Genre: "fantasy"}
		query, params := buildMovieRentalQuery(request)

		assert.Equal(t, "SELECT * FROM movies WHERE LOWER(genre) LIKE LOWER('%' || $1 || '%')", query)
		assert.Equal(t, 1, len(params))
		assert.Equal(t, "fantasy", params[0])
	})

	t.Run("create get all movies query with matching given year", func(t *testing.T) {

		request := &request.MoviesRequest{ Year: "2005"}
		query, params := buildMovieRentalQuery(request)

		assert.Equal(t, "SELECT * FROM movies WHERE LOWER(year)=$1", query)
		assert.Equal(t, 1, len(params))
		assert.Equal(t, "2005", params[0])
	})

	t.Run("create get all movies query with matching given actor", func(t *testing.T) {
		request := &request.MoviesRequest{ Actors: "robert"}
		query, params := buildMovieRentalQuery(request)

		assert.Equal(t, "SELECT * FROM movies WHERE LOWER(actors) LIKE LOWER('%' || $1 || '%')", query)
		assert.Equal(t, 1, len(params))
		assert.Equal(t, "robert", params[0])
	})

	t.Run("create get all movies query with 2 matching search queries", func(t *testing.T) {

		request := &request.MoviesRequest{ Actors: "robert", Year: "2005"}
		query, params := buildMovieRentalQuery(request)

		assert.Equal(t, "SELECT * FROM movies WHERE LOWER(year)=$1 AND LOWER(actors) LIKE LOWER('%' || $2 || '%')", query)
		assert.Equal(t, 2, len(params))
		assert.Equal(t, "2005", params[0])
		assert.Equal(t, "robert", params[1])
	})

	t.Run("create get all movies query with 3 matching search queries", func(t *testing.T) {

		request := &request.MoviesRequest{ Actors: "robert", Year: "2005", Genre: "action"}
		query, params := buildMovieRentalQuery(request)

		assert.Equal(t, "SELECT * FROM movies WHERE LOWER(genre) LIKE LOWER('%' || $1 || '%') AND LOWER(year)=$2 AND LOWER(actors) LIKE LOWER('%' || $3 || '%')", query)
		assert.Equal(t, 3, len(params))
		assert.Equal(t, "action", params[0])
		assert.Equal(t, "2005", params[1])
		assert.Equal(t, "robert", params[2])
	})
}
