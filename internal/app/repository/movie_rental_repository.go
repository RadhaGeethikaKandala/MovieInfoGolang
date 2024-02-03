package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
)

type Repository interface {
	GetMovies(query string, params []string) []dto.Movie
	GetRatingsFor(movieId int) []dto.Rating
	CreateQuery(movieRequest *request.MoviesRequest) (string, []string)
}

type repository struct {
	db *sql.DB
}

func (r *repository) CreateQuery(movieRequest *request.MoviesRequest) (string, []string) {
	query := "SELECT * FROM movies"
	param_position_count := 1
	conditional_regex_query := " LOWER(%s) LIKE LOWER('%%' || $%d || '%%')"
	conditional_exact_query := " LOWER(%s)=$%d"

	params := []string{}

	whereFlag := false

	if movieRequest.Genre != "" {
		if !whereFlag {
			query += " WHERE" + fmt.Sprintf(conditional_regex_query, "genre", param_position_count)
			whereFlag = true
		} else {
			query += " AND" + fmt.Sprintf(conditional_regex_query, "genre", param_position_count)
		}
		params = append(params, movieRequest.Genre)
		param_position_count++
	}
	if movieRequest.Year != "" {
		if !whereFlag {
			query += " WHERE" + fmt.Sprintf(conditional_exact_query, "year", param_position_count)
			whereFlag = true
		} else {
			query += " AND" + fmt.Sprintf(conditional_exact_query, "year", param_position_count)
		}
		params = append(params, movieRequest.Year)
		param_position_count++
	}
	if movieRequest.Actors != "" {
		if !whereFlag {
			query += " WHERE" + fmt.Sprintf(conditional_regex_query, "actors", param_position_count)
			whereFlag = true
		} else {
			query += " AND" + fmt.Sprintf(conditional_regex_query, "actors", param_position_count)
		}
		params = append(params, movieRequest.Actors)
		param_position_count++
	}

	return query, params
}

func (r *repository) GetMovies(query string, params []string) []dto.Movie {

	// query := `SELECT * FROM movies m
	// 					 LEFT JOIN moviesratings mr
	// 					 ON m.id = mr.movie_id
	// 					 RIGHT JOIN ratings r
	// 					 ON mr.rating_id = r.id`

	params_conv := transformParams(params)

	rows, err := r.db.Query(query, params_conv...)
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

func transformParams(params []string) []interface{} {
	var params_conv []interface{}
	for _, param := range params {
		params_conv =append(params_conv, param)
	}
	return params_conv
}

func NewRepository(conn *sql.DB) Repository {

	return &repository{
		db: conn,
	}
}
