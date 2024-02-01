package service

import (
	"testing"

	mock_client "github.com/RadhaGeethikaKandala/MovieRental/internal/app/client/mocks"
	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMovieService(t *testing.T) {

	ctrl := gomock.NewController(t)
	client := mock_client.NewMockOmdbClient(ctrl)
	service := NewMovieService(client)

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

}
