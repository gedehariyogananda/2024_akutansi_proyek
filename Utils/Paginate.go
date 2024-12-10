package Utils

import (
	"2024_akutansi_project/Models/Common"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

func GetPaginationParams(ctx *gin.Context, defaultLimit, defaultPage int) (int, int) {
	perPage, err := strconv.Atoi(ctx.DefaultQuery("limit", strconv.Itoa(defaultLimit)))
	if err != nil || perPage < 1 {
		perPage = defaultLimit
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", strconv.Itoa(defaultPage)))
	if err != nil || page < 1 {
		page = defaultPage
	}

	return perPage, page
}

func InsertParams(ctx *gin.Context) Common.Query {
	perPage, err := strconv.Atoi(ctx.DefaultQuery("limit", strconv.Itoa(Common.DEFAULTLIMIT)))
	if err != nil || perPage < 1 {
		perPage = Common.DEFAULTLIMIT
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", strconv.Itoa(Common.DEFAULTPAGE)))
	if err != nil || page < 1 {
		page = Common.DEFAULTPAGE
	}

	query := Common.Query{
		Limit: perPage,
		Page:  page,
	}

	if ctx.GetString("company_id") != "" {
		companyID := ctx.GetString("company_id")
		query.CompanyID = &companyID
	}

	if ctx.Query("search") != "" {
		search := ctx.Query("search")
		query.Search = &search
	}

	if ctx.Query("is_locked") != "" {
		isLocked, _ := strconv.ParseBool(ctx.Query("is_locked"))
		query.IsLocked = isLocked
	}

	if ctx.Query("status") != "" {
		status, _ := strconv.ParseBool(ctx.Query("status"))
		query.Status = status
	}

	return query
}

func CountModelRecords(db *gorm.DB, model interface{}) (int64, error) {
	var totalData int64
	if err := db.Model(model).Count(&totalData).Error; err != nil {
		return 0, err
	}
	return totalData, nil
}

