package request

type MoviesRequest struct {
	Genre string `json:"genre"`
	Actors string `json:"actors"`
	Year string `json:"year"`
}