package Common

type Query struct {
	Search         *string `json:"search"`
	IsLocked       *bool   `json:"is_locked"`
	Status         *bool   `json:"status"`
	Limit          int     `json:"limit"`
	TypeAccount    *string `json:"type_account"`
	Page           int     `json:"page"`
	CompanyID      *string `json:"company_id"`
	SmallestUnitID string  `json:"smallest_unit_id"`
	CategoryID     *string `json:"category_id"`
}
