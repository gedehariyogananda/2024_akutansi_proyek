package Models

type UserCompany struct {
	UserID    int
	CompanyID int
	User      User    `gorm:"foreignKey:UserID"`
	Company   Company `gorm:"foreignKey:CompanyID"`
}
