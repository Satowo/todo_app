package types

type CreateItemRequest struct {
	CategoryID uint64 `json:"category_id"`
	ItemTitle string `json:"item_title"`
	Content string `json:"content"`
	ExpiredAt string `json:"expired_at"`
}

type UpdateItemRequest struct {
	CategoryID uint64 `json:"category_id"`
	ItemTitle string `json:"item_title"`
	Content string `json:"content"`
	ExpiredAt string `json:"expired_at"`
}