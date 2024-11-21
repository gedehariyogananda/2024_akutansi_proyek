package Dto

type MakeWaitingListRequest struct {
	Name         string `json:"name" binding:"required"`
	Company      string `json:"company" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Email        string `json:"email" binding:"required"`
	BusinessType string `json:"business_type" binding:"required"`
}
