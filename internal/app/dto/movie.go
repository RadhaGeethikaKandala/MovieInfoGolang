package dto

type Movie struct {
	Title      string `json:"title"`
	Year       string `json:"year"`
	Rated      string `json:"rated"`
	Released   string `json:"released"`
	Runtime    string `json:"runtime"`
	Genre      string `json:"genre"`
	Director   string `json:"director"`
	Writer     string `json:"writer"`
	Actors     string `json:"actors"`
	Plot       string `json:"plot"`
	Language   string `json:"language"`
	Country    string `json:"country"`
	Awards     string `json:"awards"`
	Poster     string `json:"poster"`
	Ratings    []Rating `json:"ratings"`
	Metascore  string `json:"metascore"`
	ImdbRating string `json:"imdb_rating"`
	ImdbVotes  string `json:"imdb_votes"`
	ImdbID     string `json:"imdb_id"`
	Type       string `json:"type"`
	DVD        string `json:"dvd"`
	BoxOffice  string `json:"box_office"`
	Production string `json:"production"`
	Website    string `json:"website"`
	Response   string `json:"reponse"`
}
