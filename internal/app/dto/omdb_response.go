package dto

type OmdbResponse struct {
	Search   []Movie
	Response string
	Error    string
}
