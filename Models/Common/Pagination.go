package Common

import (
	"os"
	"strconv"
)

type Meta struct {
	PerPage         int    `json:"per_page"`
	Page            int    `json:"page"`
	TotalData       int64  `json:"total_data"`
	TotalPage       int64  `json:"total_page"`
	PreviousPageURL string `json:"previous_page"`
	NextPageURL     string `json:"next_page"`
}

func PaginateMetadata(totalData int64, perPage, page int, path string) (Meta Meta) {
	Meta.PerPage = perPage
	Meta.Page = page
	Meta.TotalData = totalData

	if perPage > 0 {
		Meta.TotalPage = (totalData + int64(perPage) - 1) / int64(perPage)
	} else {
		Meta.TotalPage = 0
	}

	Meta.PreviousPageURL = getPageURL(page-1, perPage, Meta.TotalPage, path)
	Meta.NextPageURL = getPageURL(page+1, perPage, Meta.TotalPage, path)

	return Meta
}

// getPageURL Paginate
func getPageURL(page int, perPage int, totalPage int64, path string) string {
	if page < 1 || page > int(totalPage) {
		return "-"
	}
	return os.Getenv("API_URL_V1") + path + "?perPage=" + strconv.Itoa(perPage) + "&page=" + strconv.Itoa(page)
}
