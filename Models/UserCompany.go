package Models

type UserCompany struct {
	UserID    int     `json:"user_id"`
	CompanyID int     `json:"company_id"`
	User      User    `gorm:"foreignKey:UserID"`
	Company   Company `gorm:"foreignKey:CompanyID"`
}
