package request

type MoviesRequest struct {
	Genre string `json:"genre" form:"genre"`
	Actors string `json:"actors" form:"actors"`
	Year string `json:"year" form:"year"`
}