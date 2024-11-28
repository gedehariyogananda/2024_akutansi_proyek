package Models

type DeviceToken struct {
	UserID      string `json:"user_id" bson:"user_id"`
	DeviceToken string `json:"device_token" bson:"device_token"`
}
