package service

import (
	"testing"

	mock_client "github.com/RadhaGeethikaKandala/MovieRental/internal/app/client/mocks"
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
	client := mock_client.NewMockOmdbClient(ctrl)
	service := NewMovieService(client, repository)

	t.Run("it should return movie list", func(t *testing.T) {
		movieName := "batman"
		movieList := []dto.Movie{
			{
				Title: "batman",
			},
			{
				Title: "batman returns",
			},
		}

		var omdbResponse = dto.OmdbResponse{
			Search: movieList,
		}

		client.EXPECT().GetMovieList(movieName).Times(1).Return(omdbResponse)
		movies, _ := service.GetMovies(movieName)

		assert.Equal(t, 2, len(movies))
		assert.Equal(t, movieName, movies[0].Title)
		assert.Equal(t, movieName+" returns", movies[1].Title)
	})

	t.Run("it should return error if movie name is not found", func(t *testing.T) {
		movieName := "xyz"
		var omdbResponse = dto.OmdbResponse{
			Response: "False",
			Error:    "Movie not found!",
		}

		client.EXPECT().GetMovieList(movieName).Times(1).Return(omdbResponse)

		movies, err := service.GetMovies(movieName)

		assert.Equal(t, 0, len(movies))
		assert.Equal(t, "Movie not found!", err.Error())
	})




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
	client := mock_client.NewMockOmdbClient(ctrl)
	service := NewMovieService(client, repository)

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
