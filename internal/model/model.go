package model

type Board struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Deleted  bool `json:"deleted"`
}

type Category struct {
	ID        uint64 `json:"id"`
	BoardID   uint64 `json:"board_id"`
	Title     string `json:"title"`
	Deleted  bool `json:"deleted"`
}

type Item struct {
	ID        uint64 `json:"id"`
	CategoryID uint64 `json:"category_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ExpiredAt string `json:"expired_at"`
	Archived  bool `json:"archived"`
}
