package repository

// import (
// 	"database/sql"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetMoviesFromDb(t *testing.T) {
// 	var db *sql.DB
// 	repository := NewRepository(db)
// 	movies := repository.GetMovies()

// 	assert.Equal(t, 1, len(movies))
// 	assert.Equal(t, "Batman Begins", movies[0].Title)
// 	assert.Equal(t, 3, len(movies[0].Ratings))
// }