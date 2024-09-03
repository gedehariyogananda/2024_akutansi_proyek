package Models

import "time"

type Company struct {
	ID        int
	Name      string
	Address   string
	ImageCompany     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
