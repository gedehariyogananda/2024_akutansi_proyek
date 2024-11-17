package Dto

type CallbackProfile struct {
	ID       string `json:"id" bson:"_id"`
	Price    string `json:"price" bson:"price"`
	Provider string `json:"provider" bson:"provider"`
}

type ShopeeCallbackRequest struct {
	ID    string `json:"id" bson:"_id"`
	Price string `json:"price"`
}
