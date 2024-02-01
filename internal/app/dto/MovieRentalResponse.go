package dto

type MovieRentalResponse struct {
	MovieList    []Movie `json:"movieList"`
	ErrorMessage string  `json:"errorMessage"`
}
