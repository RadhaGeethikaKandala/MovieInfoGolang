package service

import (
	"testing"

	mock_client "github.com/RadhaGeethikaKandala/MovieRental/internal/app/client/mocks"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto/request"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMovieService(t *testing.T) {

	ctrl := gomock.NewController(t)
	repository := mock_repository.NewMockRepository(ctrl)
	client := mock_client.NewMockOmdbClient(ctrl)
	service := NewMovieService(client, repository)

	t.Run("service should return movie list", func(t *testing.T) {
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

	t.Run("service should return error if movie name is not found", func(t *testing.T) {
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

	t.Run("service should return all movies from db", func(t *testing.T) {
		movieTestData := []dto.Movie{
			{Id: 1, Title: "batman"},
			{Id: 2, Title: "batman returns"},
		}

		ratingsTestData1 := []dto.Rating{
			{Id: 2, Source: "Rotten Tomatoes", Value: "85%"},
		}

		request := &request.MoviesRequest{}
		repository.EXPECT().GetMovies(request).Times(1).Return(movieTestData)
		repository.EXPECT().GetRatingsFor(gomock.Any()).Times(2).
			Return(ratingsTestData1)

		movies := service.GetMoviesFromDb(request)

		assert.Equal(t, 2, len(movies))
		assert.Equal(t, "batman", movies[0].Title)
		assert.Equal(t, 1, len(movies[0].Ratings))
		assert.Equal(t, "Rotten Tomatoes", movies[0].Ratings[0].Source)

		assert.Equal(t, 1, len(movies[1].Ratings))
		assert.Equal(t, "batman returns", movies[1].Title)
		assert.Equal(t, "Rotten Tomatoes", movies[0].Ratings[0].Source)

	})

	t.Run("service should return movies with matching given genre", func(t *testing.T) {
		movieTestData := []dto.Movie{

			{
				Title:  "spiderman2",
				Genre:  "Fantasy",
				Actors: "Robert",
			},
		}

		ratingsTestData1 := []dto.Rating{
			{Id: 2, Source: "Rotten Tomatoes", Value: "85%"},
		}

		request := &request.MoviesRequest{Genre: "Fantasy", Actors: "Robert"}
		repository.EXPECT().GetMovies(request).Times(1).Return(movieTestData)
		repository.EXPECT().GetRatingsFor(gomock.Any()).Times(1).
			Return(ratingsTestData1)

		movies := service.GetMoviesFromDb(request)

		assert.Equal(t, 1, len(movies))
		assert.Equal(t, "spiderman2", movies[0].Title)
		assert.Equal(t, "Fantasy", movies[0].Genre)
		assert.Equal(t, "Robert", movies[0].Actors)

	})

}
