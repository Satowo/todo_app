package types

type CreateBoardRequest struct {
	BoardTitle string `json:"board_title"`
}

type CreateCategoryRequest struct {
	CategoryTitle string `json:"category_title"`
}