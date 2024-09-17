package Models

type UserCompany struct {
	UserID    string  `json:"user_id"`
	CompanyID string  `json:"company_id"`
	User      User    `gorm:"foreignKey:UserID"`
	Company   Company `gorm:"foreignKey:CompanyID"`
}
