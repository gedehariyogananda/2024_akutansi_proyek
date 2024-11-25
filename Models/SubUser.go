package Models

import (
	"2024_akutansi_project/Utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	EmployeeKey string `json:"employee_key"`
	Password    string `json:"password"`
	Status      bool   `json:"status"`
	CompanyID   string `json:"company_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

// create uuid
func (subUser *SubUser) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid
	if subUser.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		subUser.ID = uuid.String()
	}

	if subUser.Password != "" {
		hashed, err := Utils.HashPassword(subUser.Password)
		if err != nil {
			return err
		}

		subUser.Password = hashed
	}

	return
}
