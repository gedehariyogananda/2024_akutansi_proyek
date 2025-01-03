package Dto

import "2024_akutansi_project/Models"

type CreateAccountDto struct {
	Name      string             `json:"name" binding:"required"`
	Type      Models.TypeAccount `json:"type" binding:"required,oneof=ASSET LIABILITY EQUITY REVENUE EXPENSE"`
	Code      string             `json:"code" binding:"required"`
	CompanyID string             `json:"-"`
	Status    bool               `json:"status"`
	IsLocked  bool               `json:"is_locked"`
}

type UpdateAccountDto struct {
	Name      string             `json:"name"`
	Type      Models.TypeAccount `json:"type" binding:"oneof=ASSET LIABILITY EQUITY REVENUE EXPENSE"`
	Code      string             `json:"code"`
	CompanyID string             `json:"-"`
	Status    bool               `json:"status"`
	IsLocked  bool               `json:"is_locked"`
}
