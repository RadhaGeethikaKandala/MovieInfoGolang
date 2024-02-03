package request

type MoviesRequest struct {
	ImdbId string `json:"imdbid" form:"imdbid"`
	Genre  string `json:"genre" form:"genre"`
	Actors string `json:"actors" form:"actors"`
	Year   string `json:"year" form:"year"`
}
