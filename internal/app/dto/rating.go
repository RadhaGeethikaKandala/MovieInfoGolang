package dto

type Rating struct {
	Id     int    `json:"-"`
	Source string `json:"source"`
	Value  string `json:"value"`
}
