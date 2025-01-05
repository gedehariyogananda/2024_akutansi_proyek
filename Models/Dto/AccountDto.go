package Dto

import "2024_akutansi_project/Models"

type CreateAccountDto struct {
	Name      string             `json:"name" validate:"required"`
	Type      Models.TypeAccount `json:"type" validate:"required,oneof=ASSET LIABILITY EQUITY REVENUE EXPENSE"`
	Code      string             `json:"code" validate:"required"`
	CompanyID string             `json:"-"`
	Status    bool               `json:"status"`
}

type UpdateAccountDto struct {
	Name      string             `json:"name"`
	Type      Models.TypeAccount `json:"type,omitempty" validate:"omitempty,oneof=ASSET LIABILITY EQUITY REVENUE EXPENSE"`
	Code      string             `json:"code"`
	CompanyID string             `json:"-"`
	Status    bool               `json:"status"`
}
