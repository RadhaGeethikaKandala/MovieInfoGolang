package service

import (
	"errors"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/client"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/response"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/repository"
)

type MovieService interface {
	GetMovies(movieName string) ([]dto.Movie, error)
	GetMoviesFromDb(*request.MoviesRequest) []response.TruncatedMovieResponse
	GetMovieDetails(imdbid string) dto.Movie
}

type movieService struct {
	omdbClient client.OmdbClient
	repository repository.Repository
}


func (ms *movieService) GetMovies(movieName string) ([]dto.Movie, error) {
	omdbReponse := ms.omdbClient.GetMovieList(movieName)
	movieList := omdbReponse.Search
	if apiError := omdbReponse.Error; len(apiError) > 0 {
		return nil, errors.New(apiError)
	}
	return movieList, nil
}

// To be decomissioned
// func (ms *movieService) GetMoviesFromDb(movieRequest *request.MoviesRequest) []dto.Movie {
	// 	movies := ms.repository.GetMovies(movieRequest)
	// 	for idx, movie := range movies {
		// 		movies[idx].Ratings = ms.repository.GetRatingsFor(movie.Id)
		// 	}
		// 	return movies
		// }

// GetMovieDetails implements MovieService.
func (ms *movieService) GetMovieDetails(imdbid string) dto.Movie {
	panic("unimplemented")
}


func (ms *movieService) GetMoviesFromDb(movieRequest *request.MoviesRequest) []response.TruncatedMovieResponse {
	var truncatedMovies = make([]response.TruncatedMovieResponse, 0)
	movies := ms.repository.GetMovies(movieRequest)

	for _, movie := range movies {
		truncatedMovie := response.TruncatedMovieResponse{
			Title: movie.Title, Year: movie.Year,
			Rated: movie.Rated, Actors: movie.Actors,
			Genre: movie.Genre,
		}
		truncatedMovies = append(truncatedMovies, truncatedMovie)
	}
	return truncatedMovies
}

func NewMovieService(c client.OmdbClient, r repository.Repository) MovieService {
	return &movieService{
		omdbClient: c,
		repository: r,
	}
}
