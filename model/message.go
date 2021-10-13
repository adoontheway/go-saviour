package model

type MessageListEntity struct {
	Data []MessageListData `json:"data"`
}

type MessageListData struct {
	Stranger  MessageListDataStranger `json:"stranger"`
	Count     uint32                  `json:"count"`
	Type      string                  `json:"type"`
	Id        uint32                  `json:"id"`
	Online    uint32                  `json:"online"`
	Excerpt   string                  `json:"excerpt"`
	Draft     string                  `json:"draft"`
	Createdat string                  `json:"created_at"`
	Updatedat string                  `json:"updated_at"`
}
type MessageListDataStranger struct {
	Id                uint32 `json:"id"`
	Username          string `json:"username"`
	Nickname          string `json:"nickname"`
	Avatar            string `json:"avatar"`
	Email             string `json:"email"`
	Introduction      string `json:"introduction"`
	NotificationCount uint32 `json:"notification_count"`

	Status    uint32 `json:"status"`
	Online    uint32 `json:"online"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
