package Response

import "2024_akutansi_project/Models"

type CompanyResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address *string `json:"address"`
	Image   *string `json:"image"`
}

func ToCompanyResponse(company *Models.Company) *CompanyResponse {
	return &CompanyResponse{
		ID:      company.ID,
		Name:    company.Name,
		Address: company.Address,
		Image:   company.Image,
	}
}
