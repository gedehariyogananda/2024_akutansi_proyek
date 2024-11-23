package Response

import "2024_akutansi_project/Models"

type Category struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Status bool   `json:"status"`
}

func ToCategory(category *Models.Category) *Category {
	return &Category{
		ID:     category.ID,
		Name:   category.Name,
		Code:   category.Code,
		Type:   category.Type,
		Status: category.Status,
	}
}

func ToCategorySlice(categories []*Models.Category) []*Category {
	categoryResponses := []*Category{}
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategory(category))
	}
	return categoryResponses
}
