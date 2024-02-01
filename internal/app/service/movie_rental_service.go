package service

import (
	"errors"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/client"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
)

type MovieService interface {
	GetMovies(movieName string) ([]dto.Movie, error)
}

type movieService struct {
	omdbClient client.OmdbClient
}

func (ms *movieService) GetMovies(movieName string) ([]dto.Movie, error) {
	omdbReponse := ms.omdbClient.GetMovieList(movieName)
	movieList := omdbReponse.Search
	if apiError := omdbReponse.Error; len(apiError) > 0 {
		return nil, errors.New(apiError)
	}
	return movieList, nil
}

func NewMovieService(client client.OmdbClient) MovieService {
	return &movieService{
		omdbClient: client,
	}
}
