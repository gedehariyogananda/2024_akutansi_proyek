package Models

type Tax struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Precentage int64  `json:"precentage"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
