package response

import "github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"

type MovieResponse struct {
	Movie dto.Movie `json:"movie"`
}
