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
	GetMoviesFromDb(*request.MoviesRequest) response.TruncatedMovieReponse
	GetMovieDetails(imdbid string) response.MovieResponse
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

// GetMovieDetails implements MovieService.
func (ms *movieService) GetMovieDetails(imdbid string) response.MovieResponse {
	movie := ms.repository.GetMovie(imdbid)
	movie.Ratings = ms.repository.GetRatingsFor(movie.Id)
	movieResponse := response.MovieResponse{Movie: movie}
	return movieResponse
}


func (ms *movieService) GetMoviesFromDb(movieRequest *request.MoviesRequest) response.TruncatedMovieReponse {
	var truncatedMovies = make([]response.TruncatedMovie, 0)
	movies := ms.repository.GetMovies(movieRequest)

	for _, movie := range movies {
		truncatedMovie := response.TruncatedMovie{
			Title: movie.Title, Year: movie.Year,
			Rated: movie.Rated, Actors: movie.Actors,
			Genre: movie.Genre,
		}
		truncatedMovies = append(truncatedMovies, truncatedMovie)
	}
	return response.TruncatedMovieReponse{Movies: truncatedMovies}
}

func NewMovieService(c client.OmdbClient, r repository.Repository) MovieService {
	return &movieService{
		omdbClient: c,
		repository: r,
	}
}
