package Common

type Query struct {
	Search   *string `json:"serch"`
	IsLocked bool    `json:"is_locked"`
	Status   bool    `json:"status"`
	Limit    int     `json:"limit"`
	Page     int     `json:"page"`
}
