package model

type DynamicCommentEntity struct {
	Data  []DynamicCommentData `json:"data"`
	Links DynamicCommentLink   `json:"links"`
	Meta  DynamicCommentMeta   `json:"meta"`
}

type DynamicCommentData struct {
	Id               uint32 `json:"id"`
	ReplierId        uint32 `json:"replier_id"`
	PosterId         uint32 `json:"poster_id"`
	ReplierNicknam   string `json:"replier_nickname"`
	ReplierAvatar    string `json:"replier_avatar"`
	ReceiverId       uint32 `json:"receiver_id"`
	ReceiverNickname string `json:"receiver_nickname"`
	ThumbCount       uint32 `json:"thumb_count"`
	Content          string `json:"content"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type DynamicCommentLink struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type DynamicCommentMeta struct {
	CurrentPage uint32                   `json:"current_page"`
	From        uint32                   `json:"from"`
	LastPage    uint32                   `json:"last_page"`
	Path        string                   `json:"path"`
	PerPage     string                   `json:"per_page"`
	To          uint32                   `json:"to"`
	Total       uint32                   `json:"total"`
	Links       []DynamicCommentMetaLink `json:"link"`
}

type DynamicCommentMetaLink struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
