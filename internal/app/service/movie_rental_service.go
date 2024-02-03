package service

import (
	"errors"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/client"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/repository"
)

type MovieService interface {
	GetMovies(movieName string) ([]dto.Movie, error)
	GetMoviesFromDb(*request.MoviesRequest) []dto.Movie
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

func (ms *movieService) GetMoviesFromDb(movieRequest *request.MoviesRequest) []dto.Movie {
	db_query_url, params := ms.repository.CreateQuery(movieRequest)
	movies := ms.repository.GetMovies(db_query_url, params)
	for idx, movie := range movies {
		movies[idx].Ratings = ms.repository.GetRatingsFor(movie.Id)
	}
	return movies
}

func NewMovieService(c client.OmdbClient, r repository.Repository) MovieService {
	return &movieService{
		omdbClient: c,
		repository: r,
	}
}
