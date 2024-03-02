package model

type Board struct {
	ID        int `json:"id"`
	Title     string `json:"title"`
	Deleted  bool `json:"deleted"`
}

type Category struct {
	ID        int `json:"id"`
	BoardID   int `json:"board_id"`
	Title     string `json:"title"`
	Deleted  bool `json:"deleted"`
}

type Item struct {
	ID        int `json:"id"`
	CategoryID int `json:"category_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ExpiredAt string `json:"expired_at"`
	Deleted  bool `json:"deleted"`
}
