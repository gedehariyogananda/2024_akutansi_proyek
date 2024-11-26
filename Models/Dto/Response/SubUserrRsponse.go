package Response

import "2024_akutansi_project/Models"

type SubUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	EmployeeKey string `json:"employee_key"`
	Status      bool   `json:"status"`
	CompanyID   string `json:"company_id"`
}

func ToSubUser(subUser *Models.SubUser) *SubUser {
	return &SubUser{
		ID:          subUser.ID,
		Name:        subUser.Name,
		EmployeeKey: subUser.EmployeeKey,
		Status:      subUser.Status,
		CompanyID:   subUser.CompanyID,
	}
}

func ToSubUserSlice(subUsers []*Models.SubUser) []*SubUser {
	subUserResponses := []*SubUser{}
	for _, subUser := range subUsers {
		subUserResponses = append(subUserResponses, ToSubUser(subUser))
	}
	return subUserResponses
}
