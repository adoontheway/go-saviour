package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Collection struct {
	DynamicID int64  `bson:"dynamic_id,omitempty"`
	UserId    int64  `bson:"user_id,omitempty"`
	VideoId   int64  `bson:"video_id,omitempty"`
	Type      string `bson:"type,omitempty"`
}

type DComment struct {
	DynamicID        int64  `bson:"dynamic_id,omitempty"`
	ReceiverID       int64  `bson:"receiver_id,omitempty"`
	Content          string `bson:"content,omitempty"`
	ThumbCount       int64  `bson:"thumb_count,omitempty"`
	ReceiverNickname string `bson:"receiver_nickname,omitempty"`
	ReplierID        int64  `bson:"replier_id,omitempty"`
	ReplierNickname  string `bson:"replier_nickname,omitempty"`
	ReplierAvatar    string `bson:"replier_avatar,omitempty"`
	PosterId         int    `bson:"poster_id,omitempty"`
}

type Dynamic struct {
	ID           primitive.ObjectID `bson:"_id,omiepty"`
	UserId       int64              `bson:"user_id,omitempty"`
	Content      string             `bson:"content,omitempty"`
	ThumbCount   int64              `bson:"thumb_count,omitempty"`
	CollectCount int64              `bson:"collect_count,omitempty"`
	CommentCount int64              `bson:"comment_count,omitempty"`
	Status       string             `bson:"status,omitempty"`
	Location     string             `bson:"location,omitempty"`
}

type Friend struct {
}

type Image struct {
	ID        primitive.ObjectID `bson:"_id,omiepty"`
	Type      string             `bson:"type,omitempty"`
	Path      string             `bson:"path,omitempty"`
	UserID    int64              `bson:"user_id,omitempty"`
	DynamicID int64              `bson:"dynamic_id,omitempty"`
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omiepty"`
	FromID    string             `bson:"from_id,omitempty"`
	ToID      string             `bson:"to_id,omitempty"`
	Content   string             `bson:"content,omitempty"`
	Push      int64              `bson:"push,omitempty"`
	DynamicID int64              `bson:"dynamic_id,omitempty"`
	VideoId   int64              `bson:"vlog_id,omitempty"`
	Type      string             `bson:"type,omitempty"` // vlog system
	Read      int64              `bson:"read,omitempty"`
}

type Record struct {
	ID        primitive.ObjectID `bson:"_id,omiepty"`
	FromID    string             `bson:"from_id,omitempty"`
	ToID      string             `bson:"to_id,omitempty"`
	Content   string             `bson:"content,omitempty"`
	Push      int64              `bson:"push,omitempty"`
	Status    int64              `bson:"status,omitempty"`
	UUID      string             `bson:"uuid,omitempty"`
	Type      string             `bson:"type,omitempty"` // vlog system
	Read      int64              `bson:"read,omitempty"`
	ThumbNail string             `bson:"thumbnail,omitempty"` // vlog system
	Length    int64              `bson:"length,omitempty"`
}

type Stranger struct {
	UserId     int64 `bson:"user_id,omiemty"`
	StrangerID int64 `bson:"stranger_id,omiemty"`
	Hide       bool  `bson:"hide,omiemty"`
}

type Thumb struct {
	Type      string `bson:"type,omitempty"` // dynamic
	VideoId   int64  `bson:"vlog_id,omitempty"`
	DynamicID int64  `bson:"dynamic_id,omitempty"`
	UserID    int64  `bson:"user_id,omitempty"`
}

type User struct {
	ID                 primitive.ObjectID `bson:"_id,omiepty"`
	UserName           string             `bson:"username,omitempty"`
	NickName           string             `bson:"nickname,omitempty"`
	Password           string             `bson:"password,omitempty"`
	Avatar             int64              `bson:"avatar"`
	Phone              int64              `bson:"phone,omitempty"`
	Email              string             `bson:"email,omitempty"`
	EmailVerifyAT      string             `bson:"email_verify_at,omitempty"` // vlog system
	Introduction       int64              `bson:"introduction,omitempty"`
	Provifer           string             `bson:"provider,omitempty"` // vlog system
	ProviderID         int64              `bson:"provider_id,omitempty"`
	NotificationCounts int64              `bson:"notification_count,omitempty"`
	Status             string             `bson:"status,omitempty"` // vlog system
	Online             int64              `bson:"online,omitempty"`
	CreatedAt          string             `bson:"status,omitempty"` // vlog system
	UpdatedAt          int64              `bson:"online,omitempty"`
}

type VLog struct {
	ID           primitive.ObjectID `bson:"_id,omiepty"`
	UserID       int64              `bson:"user_id,omiempty"`
	ThumbCount   int64              `bson:"thumb_count,omitempty"`
	CollectCount int64              `bson:"collect_count,omitempty"`
	CommentCount int64              `bson:"comment_count,omitempty"`
	Illustration int64              `bson:"illustration"`
	Status       int64              `bson:"status,omitempty"`
	Location     string             `bson:"location,omitempty"`
	Title        string             `bson:"title,omitempty"` // vlog system

}

type Video struct {
	Type    string `bson:"type,omitempty"` // vlog
	Path    string `bson:"path,omitempty"`
	VideoID int64  `bson:"vlog_id,omitempty"`
	UserID  int64  `bson:"user_id,omitempty"`
}

type VComment struct {
	VideoID    int64  `bson:"vlog_id,omitempty"`
	UserID     int64  `bson:"user_id,omitempty"`
	Content    string `bson:"content,omitempty"`
	ThumbCount int64  `bson:"thumb_count,omitempty"`
	Status     int64  `bson:"status,omitempty"`
	Type       string `bson:"type,omitempty"` // dynamic
}
