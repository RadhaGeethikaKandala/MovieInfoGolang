package repository

import (
	"fmt"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
)

func buildMovieRentalQuery(movieRequest *request.MoviesRequest) (string, []interface{}) {
	query := "SELECT * FROM movies"
	conditional_regex_query := " LOWER(%s) LIKE LOWER('%%' || $%d || '%%')"
	conditional_exact_query := " LOWER(%s)=$%d"

	params := []string{}
	param_position_count := 1
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

	return query, transformParams(params)
}

func transformParams(params []string) (params_conv []interface{}) {
	for _, param := range params {
		params_conv = append(params_conv, param)
	}
	return
}
