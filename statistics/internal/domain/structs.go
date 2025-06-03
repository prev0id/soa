package domain

type Event struct {
	PromoID uint64 `json:"promo_id"`
	PostID  uint64 `json:"post_id"`
	UserID  uint64 `json:"user_id"`
}

type PromoStats struct {
	Views    uint64
	Likes    uint64
	Comments uint64
}
