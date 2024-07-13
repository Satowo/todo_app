package types

type CreateBoardRequest struct {
	BoardTitle string `json:"board_title"`
}

type UpdateBoardRequest struct {
	BoardTitle string `json:"board_title"`
}