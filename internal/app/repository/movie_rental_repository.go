package repository

import (
	"database/sql"
	"log"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
)

type Repository interface {
	GetMovies() []dto.Movie
	GetRatingsFor(movieId int) []dto.Rating
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetMovies() []dto.Movie {

	// query := `SELECT * FROM movies m
	// 					 LEFT JOIN moviesratings mr
	// 					 ON m.id = mr.movie_id
	// 					 RIGHT JOIN ratings r
	// 					 ON mr.rating_id = r.id`

	query := "SELECT * FROM movies"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatalf("Error while querying data: %s", err.Error())
	}
	defer rows.Close()

	movies := make([]dto.Movie, 0)
	for rows.Next() {
		var movie dto.Movie
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Rated, &movie.Released,
			&movie.Runtime, &movie.Genre, &movie.Director, &movie.Writer,
			&movie.Actors, &movie.Plot, &movie.Language, &movie.Country, &movie.Awards,
			&movie.Poster, &movie.Metascore, &movie.ImdbRating, &movie.ImdbVotes, &movie.ImdbID,
			&movie.Type, &movie.DVD, &movie.BoxOffice, &movie.Production, &movie.Website, &movie.Response)

		if err != nil {
			log.Fatalf("Error while scaning data: %s", err.Error())
		}
		movies = append(movies, movie)
	}
	return movies
}


func (r *repository) GetRatingsFor(movieId int) []dto.Rating {
	query := "SELECT * FROM ratings WHERE id IN (SELECT rating_id FROM moviesratings WHERE movie_id=$1)"
	rows, err := r.db.Query(query, movieId)
	if err != nil {
		log.Fatalf("[GetRatingsFor] Error while querying data: %s", err.Error())
	}
	defer rows.Close()
	ratings := make([]dto.Rating, 0)

	for rows.Next() {
		var rating dto.Rating
		err = rows.Scan(&rating.Id, &rating.Source, &rating.Value)

		if err != nil {
			log.Fatalf("[GetRatingsFor] Error while scaning data: %s", err.Error())
		}
		ratings = append(ratings, rating)
	}

	return ratings
}


func NewRepository(conn *sql.DB) Repository {

	return &repository{
		db: conn,
	}
}
