package types

type CreateCategoryRequest struct {
	BoardID uint64 `json:"board_id"`
	CategoryTitle string `json:"category_title"`
}

type UpdateCategoryRequest struct {
	CategoryTitle string `json:"category_title"`
}