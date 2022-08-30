package model

type Order struct {
	UID             string    `json:"order_uid"`
	TrackNumber     string    `json:"track_number"`
	Entry           string    `json:"entry"`
	Delivery        *Delivery `json:"delivery"`
	Payment         *Payment  `json:"payment"`
	Items           []*Item   `json:"items"`
	Locale          string    `json:"locale"`
	Signature       string    `json:"internal_signature"`
	CustomerID      string    `json:"customer_id"`
	DeliveryService string    `json:"delivery_service"`
	Shardkey        string    `json:"shardkey"`
	SMID            uint      `json:"sm_id"`
	DateCreated     string    `json:"date_created"`
	OffShard        string    `json:"oof_shard"`
}

func NewOrder() *Order {
	return &Order{}
}
