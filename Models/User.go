package Models

import (
	"2024_akutansi_project/Utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CompanyID string    `json:"company_id"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}

		u.ID = uuid.String()
	}

	if u.Password != "" {
		hashed, err := Utils.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = hashed
	}

	return

}
