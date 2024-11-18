package Common

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Limit           int    `json:"limit"`
	Page            int    `json:"page"`
	TotalData       int64  `json:"total_data"`
	TotalPage       int64  `json:"total_page"`
	PreviousPageURL string `json:"previous_page"`
	NextPageURL     string `json:"next_page"`
}

func PaginateMetadata(ctx *gin.Context, totalData int64, limit int, page int) Meta {
	totalPage := int64(0)
	if limit > 0 {
		totalPage = (totalData + int64(limit) - 1) / int64(limit)
	}

	path := ctx.Request.URL.Path

	meta := Meta{
		Limit:     limit,
		Page:      page,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	meta.PreviousPageURL = getPageURL(page-1, limit, totalPage, path)
	meta.NextPageURL = getPageURL(page+1, limit, totalPage, path)

	return meta
}

func getPageURL(page int, limit int, totalPage int64, path string) string {
	if page < 1 || page > int(totalPage) {
		return "-"
	}
	return os.Getenv("API_URL_V1") + path + "?limit=" + strconv.Itoa(limit) + "&page=" + strconv.Itoa(page)
}
