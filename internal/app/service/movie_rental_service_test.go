package service

import (
	"testing"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMoviesFromDb(t *testing.T) {

	ctrl := gomock.NewController(t)
	repository := mock_repository.NewMockRepository(ctrl)
	service := NewMovieService(repository)


	t.Run("it should return all movies from db", func(t *testing.T) {
		movieTestData := []dto.Movie{
			{Id: 1, Title: "batman"},
			{Id: 2, Title: "batman returns"},
		}

		request := &request.MoviesRequest{}
		repository.EXPECT().GetMovies(request).Times(1).Return(movieTestData)

		truncatedMovieReponse := service.GetMoviesFromDb(request)

		assert.Equal(t, 2, len(truncatedMovieReponse.Movies))
		assert.Equal(t, "batman", truncatedMovieReponse.Movies[0].Title)
		assert.Equal(t, "batman returns", truncatedMovieReponse.Movies[1].Title)

	})

	t.Run("it should return movies with matching given genre", func(t *testing.T) {
		movieTestData := []dto.Movie{

			{
				Title:  "spiderman2",
				Genre:  "Fantasy",
				Actors: "Robert",
			},
		}

		request := &request.MoviesRequest{Genre: "Fantasy", Actors: "Robert"}
		repository.EXPECT().GetMovies(request).Times(1).Return(movieTestData)

		truncatedMovieReponse := service.GetMoviesFromDb(request)

		assert.Equal(t, 1, len(truncatedMovieReponse.Movies))
		assert.Equal(t, "spiderman2", truncatedMovieReponse.Movies[0].Title)
		assert.Equal(t, "Fantasy", truncatedMovieReponse.Movies[0].Genre)
		assert.Equal(t, "Robert", truncatedMovieReponse.Movies[0].Actors)

	})

}

func TestGetMovieDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := mock_repository.NewMockRepository(ctrl)
	service := NewMovieService(repository)

	t.Run("it should get the entire movie details with valid imdbid", func(t *testing.T) {
		testMovieData := dto.Movie{
				Id: 1,
				Title:  "Batman returns",
				Genre:  "Fantasy",
				Actors: "Robert",
				Year: "2022",
				ImdbID: "1234",
		}

		ratingsTestData := []dto.Rating{
				{Id: 2, Source: "Rotten Tomatoes", Value: "85%"},
			}


		repository.EXPECT().GetMovie("1234").Times(1).Return(testMovieData)
		repository.EXPECT().GetRatingsFor(1).Times(1).Return(ratingsTestData)

		movieReponse, err := service.GetMovieDetails("1234")
		require.NoError(t, err)

		assert.Equal(t, "Batman returns", movieReponse.Movie.Title)
		assert.Equal(t, "2022", movieReponse.Movie.Year)
		assert.Equal(t, "Rotten Tomatoes", movieReponse.Movie.Ratings[0].Source)
		assert.Equal(t, "85%", movieReponse.Movie.Ratings[0].Value)

	})

	t.Run("it should return error response with invalid imdbid", func(t *testing.T) {
		repository.EXPECT().GetMovie("1234-invalid").Times(1).Return(dto.Movie{})

		_, err := service.GetMovieDetails("1234-invalid")

		assert.Equal(t, "no movies found with the given imdbid", err.Error())
	})
}


func TestAddMovieToCart(t *testing.T) {

	ctrl := gomock.NewController(t)
	repository := mock_repository.NewMockRepository(ctrl)
	service := NewMovieService(repository)

	t.Run("it should not throw error if movie successfully added and customer id/movie id is valid ", func(t *testing.T) {
		request := request.AddToCartRequest{
			UserId: "2",
			ImdbId: "1",
		}
		movie := dto.Movie{
			Id: 1,
			Title: "movie title",
		}
		customer := dto.Customer{
			Id: 2,
			Name: "Rahul",
			Email: "rahul@gmail.com",
		}
		repository.EXPECT().GetMovie(request.ImdbId).Times(1).Return(movie)
		repository.EXPECT().GetCustomer(request.UserId).Times(1).Return(customer)
		repository.EXPECT().AddMovieToCart(request.UserId,  request.ImdbId).Times(1).Return(nil)

		err := service.AddMovieToCart(&request)
		require.NoError(t, err)
	})

	t.Run("it should throw error if customer id is invalid ", func(t *testing.T) {
		request := request.AddToCartRequest{
			UserId: "1",
			ImdbId: "1",
		}
		movie := dto.Movie{
			Id: 1,
			Title: "movie title",
		}
		customer := dto.Customer{}

		repository.EXPECT().GetMovie(request.ImdbId).Times(1).Return(movie)
		repository.EXPECT().GetCustomer(request.UserId).Times(1).Return(customer)

		err := service.AddMovieToCart(&request)
		require.Error(t, err)

		assert.Equal(t, "no customer found with the given imdbid", err.Error())
	})

	t.Run("it should throw error if movie id is invalid ", func(t *testing.T) {
		request := request.AddToCartRequest{
			UserId: "1",
			ImdbId: "1",
		}
		movie := dto.Movie{}

		repository.EXPECT().GetMovie(request.ImdbId).Times(1).Return(movie)

		err := service.AddMovieToCart(&request)
		require.Error(t, err)

		assert.Equal(t, "no movie found with the given imdbid", err.Error())
	})
}
