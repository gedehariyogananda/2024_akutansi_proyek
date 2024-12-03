package Utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// -------------------- utils to pagination function --------------- //
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

// --------------- end function handler pagination ----------------- //
