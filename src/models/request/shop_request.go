package request

type ShopCreateRequest struct {
	ID     string `json:"id"`
	Name  string `json:"name"`
	Introduction string `json:"introduction"`
	Location  string    `json:"location"`
	Level int `json:"level"`
	Score float64 `json:"score"`
}

type ShopUpdateRequest struct {
	ID     string `json:"id"`
	Name  string `json:"name"`
	Introduction string `json:"introduction"`
	Location  string    `json:"location"`
	Level int `json:"level"`
	Score float64 `json:"score"`
}