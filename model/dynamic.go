package model

type DynamicListEntity struct {
	Data  []DynamicListData `json:"data"`
	Links DynamicListLink   `json:"links"`
	Meta  DynamicListMeta   `json:"meta"`
}

type DynamicListData struct {
	Id           uint32                 `json:"id"`
	UserId       uint32                 `json:"user_id"`
	ThumbCount   uint32                 `json:"thumb_count"`
	Thumbed      uint32                 `json:"thumbed"`
	Collected    uint32                 `json:"collected"`
	CollectCount uint32                 `json:"collect_count"`
	CommentCount uint32                 `json:"comment_count"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	Content      string                 `json:"content"`
	Location     string                 `json:"location"`
	User         DynamicListDataUser    `json:"user"`
	Images       []DynamicListDataImage `json:"images"`
}

type DynamicListLink struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type DynamicListMeta struct {
	CurrentPage uint32                `json:"current_page"`
	From        uint32                `json:"from"`
	LastPage    uint32                `json:"last_page"`
	Links       []DynamicListMetaLink `json:"links"`
	Path        string                `json:"path"`
	PerPage     string                `json:"per_page"`
	To          uint32                `json:"to"`
	Total       uint32                `json:"total"`
}

type DynamicListDataUser struct {
	Id                uint32 `json:"id"`
	Username          string `json:"username"`
	Nickname          string `json:"nickname"`
	Avatar            string `json:"avatar"`
	Email             string `json:"email"`
	Introduction      string `json:"introduction"`
	NotificationCount uint32 `json:"notification_count"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type DynamicListDataImage struct {
	Id        uint32 `json:"id"`
	UserId    uint32 `json:"user_id"`
	DynamicId uint32 `json:"dynamic_id"`
	Type      string `json:"type"`
	Path      string `json:"path"`
}

type DynamicListMetaLink struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

type DynamicDetailEntity struct {
	Data DynamicDetailData `json:"data"`
}

type DynamicDetailData struct {
	Id           uint32                   `json:"id"`
	UserId       uint32                   `json:"user_id"`
	ThumbCount   uint32                   `json:"thumb_count"`
	Thumbed      uint32                   `json:"thumbed"`
	Collected    uint32                   `json:"collected"`
	CollectCount uint32                   `json:"collect_count"`
	CommentCount uint32                   `json:"comment_count"`
	CreatedAt    string                   `json:"created_at"`
	UpdatedAt    string                   `json:"updated_at"`
	Content      string                   `json:"content"`
	Location     string                   `json:"location"`
	User         DynamicDetailDataUser    `json:"user"`
	Images       []DynamicDetailDataImage `json:"images"`
}

type DynamicDetailDataUser struct {
	Id                uint32 `json:"id"`
	Username          string `json:"username"`
	Nickname          string `json:"nickname"`
	Avatar            string `json:"avatar"`
	Email             string `json:"email"`
	Introduction      string `json:"introduction"`
	NotificationCount uint32 `json:"notification_count"`
	Status            uint8  `json:"status"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type DynamicDetailDataImage struct {
	Id        uint32 `json:"id"`
	UserId    uint32 `json:"user_id"`
	DynamicId uint32 `json:"dynamic_id"`
	Type      string `json:"type"`
	Path      string `json:"path"`
}
