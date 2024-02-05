package service

import (
	"errors"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/response"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/repository"
)

type MovieService interface {
	GetMoviesFromDb(*request.MoviesRequest) response.TruncatedMovieReponse
	GetMovieDetails(imdbid string) (response.MovieResponse, error)
	AddMovieToCart(*request.AddToCartRequest) error
}

type movieService struct {
	repository repository.Repository
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
			Genre:  movie.Genre,
			Poster: movie.Poster,
			ImdbId: movie.ImdbID,
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
		return errors.New("no customer found with the given custId")
	}
	err := ms.repository.AddMovieToCart(customer.Id, movie.Id)
	return err
}

func NewMovieService(r repository.Repository) MovieService {
	return &movieService{
		repository: r,
	}
}
