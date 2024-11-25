package Common

type Query struct {
	Search *string `json:"serch"`
	Status bool    `json:"status"`
	Limit  int     `json:"limit"`
	Page   int     `json:"page"`
}
