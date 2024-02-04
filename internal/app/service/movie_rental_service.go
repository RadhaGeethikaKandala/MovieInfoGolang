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
	GetMovieDetails(imdbid string) (response.MovieResponse, error)
	AddMovieToCart(*request.AddToCartRequest) error
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

func (ms *movieService) GetMovieDetails(imdbid string) (response.MovieResponse, error) {
	movie := ms.repository.GetMovie(imdbid)
	if movie.Id == 0 {
		return response.MovieResponse{}, errors.New("no movies found with the given imdbid")
	}
	movie.Ratings = ms.repository.GetRatingsFor(movie.Id)
	return response.MovieResponse{Movie: movie}, nil
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

func (ms *movieService) AddMovieToCart(request *request.AddToCartRequest) error {
	movie := ms.repository.GetMovie(request.ImdbId)
	if movie.Id == 0 {
		return errors.New("no movie found with the given imdbid")
	}
	customer := ms.repository.GetCustomer(request.UserId)
	if customer.Id == 0 {
		return errors.New("no customer found with the given imdbid")
	}
	err := ms.repository.AddMovieToCart(request.UserId, request.ImdbId)
	return err
}

func NewMovieService(c client.OmdbClient, r repository.Repository) MovieService {
	return &movieService{
		omdbClient: c,
		repository: r,
	}
}
