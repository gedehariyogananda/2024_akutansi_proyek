package Response

type CompanyResponseDTO struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	ImageCompany string `json:"image_company"`
}