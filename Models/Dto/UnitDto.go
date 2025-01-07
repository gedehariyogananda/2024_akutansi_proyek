package Dto

type CreateUnitDto struct {
	Name      string `json:"name" validate:"required"`
	Code      string `json:"code" validate:"required"`
	Status    bool   `json:"status"`
	CompanyID string `json:"-"`
}

type UpdateUnitDto struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    bool   `json:"status"`
	CompanyID string `json:"-"`
}
