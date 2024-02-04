package dto

type Customer struct {
	Id    int `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
