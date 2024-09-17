package Response

type CompanyResponseDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	ImageCompany string `json:"image_company"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
