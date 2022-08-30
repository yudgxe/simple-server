package model

type Item struct {
	ChrtID      uint    `json:"chrt_id"`
	TrackNumber string  `json:"track_number"`
	Price       float32 `json:"price"`
	Rid         string  `json:"rid"`
	Name        string  `json:"name"`
	Sale        float32 `json:"sale"`
	Size        string  `json:"size"`
	TotalPrice  float32 `json:"total_price"`
	NMID        uint    `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      uint    `json:"status"`
}
