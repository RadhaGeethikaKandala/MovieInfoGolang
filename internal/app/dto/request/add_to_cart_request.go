package request

type AddToCartRequest struct {
	UserId string `json:"userid"`
	ImdbId string `json:"imdbid"`
}