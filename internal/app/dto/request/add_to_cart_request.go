package request

type AddToCartRequest struct {
	Userid string `json:"userid"`
	Movieid string `json:"movieid"`
}