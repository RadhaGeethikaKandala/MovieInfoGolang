package response

type TruncatedMovie struct {
	Title      string   `json:"title"`
	Year       string   `json:"year"`
	Rated      string   `json:"rated"`
	Actors     string   `json:"actors"`
	Genre      string   `json:"genre"`
}

type TruncatedMovieReponse struct {
	Movies []TruncatedMovie `json:"movies"`
}
