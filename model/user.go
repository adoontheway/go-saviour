package model

import "time"

const (
	SEXWOMEN = "W"
	SEXMEN   = "M"
	//
	SEXUNKNOW = "U"
)

// type User struct {
// 	//用户ID
// 	Id     int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
// 	Mobile string `xorm:"varchar(20)" form:"mobile" json:"mobile"` //手机号
// 	// 用户密码= plainpwd+salt, MD5
// 	Passwd   string `xorm:"varchar(40)" form:"passwd" json:"-"` // 什么角色
// 	Avatar   string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
// 	Sex      string `xorm:"varchar(2)" form:"sex" json:"sex"`            // 什么角色
// 	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"` // 什么角色
// 	//加盐随机字符串6
// 	Salt   string `xorm:"varchar(10)" form:"salt" json:"-"`    // 什么角色
// 	Online int    `xorm:"int(10)" form:"online" json:"online"` //是否在线
// 	//前端鉴权因子,
// 	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`    // 什么角色
// 	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`     // 什么角色
// 	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"` // 什么角色
// }

type User struct {
	Id                uint32    `json:"id" xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Username          string    `json:"username" xorm:"varchar(20)" form:"username"`
	Nickname          string    `json:"nickname" xorm:"varchar(20)" form:"nickname"`
	Password          string    `json:"-" xorm:"varchar(40)" form:"passwd"`
	Avatar            string    `json:"avatar" xorm:"varchar(150)" form:"avatar"`
	Phone             string    `json:"phone" xorm:"varchar(20)" json:"mobile"`
	Email             string    `json:"email" xorm:"varchar(60)" form:"email"`
	EmailVerifiedAt   string    `json:"-" xorm:"datetime" form:"emial_verified_at"`
	Introduction      string    `json:"introduction" xorm:"varchar(150)" form:"intro"`
	Provider          string    `json:"-" xorm:"varchar(20)" form:"provider"`
	ProviderId        uint32    `json:"-" xorm:"bigint(20)" form:"provider_id"`
	NotificationCount uint32    `json:"notification_count"`
	Status            uint8     `json:"status"`
	Online            string    `json:"online"`
	Token             string    `json:"token" xorm:"varchar(40)" form:"token"`
	CreatedAt         time.Time `json:"created_at" xorm:"datetime" form:"createat"`
	UpdatedAt         time.Time `json:"updated_at" xorm:"datetime" form:"updateat"`
	Sex               string    `xorm:"varchar(2)" form:"sex" json:"sex"`
	Salt              string    `xorm:"varchar(10)" form:"salt" json:"-"`
}
