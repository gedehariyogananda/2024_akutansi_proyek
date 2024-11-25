package Response

import "2024_akutansi_project/Models"

type Unit struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    bool   `json:"status"`
	CompanyID string `json:"company_id"`
}

func ToUnit(unit *Models.Unit) *Unit {
	return &Unit{
		ID:        unit.ID,
		Name:      unit.Name,
		Code:      unit.Code,
		Status:    unit.Status,
		CompanyID: unit.CompanyID,
	}
}

func ToUnitSlice(units []*Models.Unit) []*Unit {
	unitResponses := []*Unit{}
	for _, unit := range units {
		unitResponses = append(unitResponses, ToUnit(unit))
	}
	return unitResponses
}
