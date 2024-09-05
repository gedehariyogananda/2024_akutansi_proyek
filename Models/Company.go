package Models

import "time"

type Company struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	ImageCompany string    `json:"image_company"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"updated_at"`
}
