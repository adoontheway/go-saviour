package mock

import (
	"coward-saviour/model"
	"coward-saviour/util"
	"fmt"
	"math/rand"
	"time"
)

var (
	me            *model.User
	dynamicList   *model.DynamicListEntity
	dcomments     *model.DynamicCommentEntity
	dynamicDetail *model.DynamicDetailEntity
	messageList   *model.MessageListEntity
)

type MockCenter struct {
}

func (m *MockCenter) Init() {
	localAddr := util.GetIp4Addr()
	me = &model.User{
		Id:           1,
		Username:     "test001",
		Nickname:     "test001",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Email:        "123@163.com",
		Phone:        "13500000000",
		Status:       0,
		Introduction: "我是一条酸菜鱼，又酸又菜又多余...",
		Token:        "12345678",
		Avatar:       fmt.Sprintf("http://%s:8080/assets/header01.png", localAddr),
	}

	dynamicList = &model.DynamicListEntity{
		Data: []model.DynamicListData{
			model.DynamicListData{
				Id:           10001,
				UserId:       1,
				ThumbCount:   (rand.Uint32()%1000 + 1000),
				Thumbed:      (rand.Uint32()%1000 + 1000),
				Collected:    (rand.Uint32())%1000 + 500,
				CollectCount: rand.Uint32() % 500,
				CommentCount: rand.Uint32() % 10000,
				CreatedAt:    time.Now().Format("2006-01-02 19:50:00"),
				UpdatedAt:    time.Now().Format("2006-01-02 19:50:00"),
				Content:      "it is a new day,let's fuck",
				Location:     "Shenzhen,China",
				User: model.DynamicListDataUser{
					Id:                2,
					Username:          "text002",
					Nickname:          "test002",
					Avatar:            fmt.Sprintf("http://%s:8080/assets/header01.png", localAddr),
					Email:             "133@163.com",
					Introduction:      "hehehehehhehe....",
					NotificationCount: 10,
					CreatedAt:         time.Now().Format("2006-01-02 19:50:00"),
					UpdatedAt:         time.Now().Format("2006-01-02 19:50:00"),
				},
				Images: []model.DynamicListDataImage{
					model.DynamicListDataImage{
						Id:        1,
						UserId:    2,
						DynamicId: 10001,
						Type:      "image",
						Path:      fmt.Sprintf("http://%s:8080/assets/header01.png", localAddr),
					},
				},
			},
		},
		Links: model.DynamicListLink{
			First: fmt.Sprintf("http://%s:8080/assets/header01.png", localAddr),
			Last:  fmt.Sprintf("http://%s:8080/assets/header01.png", localAddr),
			Next:  fmt.Sprintf("http://%s:8080/assets/header01.png", localAddr),
		},
		Meta: model.DynamicListMeta{
			CurrentPage: 1,
			From:        1,
			LastPage:    1,
			Path:        fmt.Sprintf("http://%s:8080/assets/header01.png", localAddr),
			PerPage:     "123",
			To:          2,
			Total:       10,
			Links: []model.DynamicListMetaLink{
				model.DynamicListMetaLink{
					Url:    fmt.Sprintf("http://%s:8080/assets/header01.png", localAddr),
					Label:  "beauty",
					Active: true,
				},
			},
		},
	}

	dcomments = &model.DynamicCommentEntity{
		Data: []model.DynamicCommentData{
			model.DynamicCommentData{
				Id:               100001,
				ReplierId:        3,
				ReplierNicknam:   "张三",
				ReplierAvatar:    fmt.Sprintf("http://%s:8080/assets/header02.png", localAddr),
				ReceiverId:       4,
				ReceiverNickname: "nidaye",
				ThumbCount:       rand.Uint32() % 9999,
				Content:          "压力巨大的又一天",
				CreatedAt:        time.Now().Format("2006-01-02 19:50:00"),
				UpdatedAt:        time.Now().Format("2006-01-02 19:50:00"),
			},
			model.DynamicCommentData{
				Id:               100002,
				ReplierId:        3,
				ReplierNicknam:   "里斯",
				ReplierAvatar:    fmt.Sprintf("http://%s:8080/assets/header04.png", localAddr),
				ReceiverId:       4,
				ReceiverNickname: "nidaye",
				ThumbCount:       rand.Uint32() % 9999,
				Content:          "呵呵",
				CreatedAt:        time.Now().Format("2006-01-02 19:50:00"),
				UpdatedAt:        time.Now().Format("2013-01-02 19:50:00"),
			},
		},
		Links: model.DynamicCommentLink{
			First: fmt.Sprintf("http://%s:8080/assets/post/post01.jpg", localAddr),
			Last:  fmt.Sprintf("http://%s:8080/assets/post/post04.jpg", localAddr),
			Prev:  fmt.Sprintf("http://%s:8080/assets/post/post02.jpg", localAddr),
			Next:  fmt.Sprintf("http://%s:8080/assets/post/post03.jpg", localAddr),
		},
		Meta: model.DynamicCommentMeta{
			CurrentPage: 1,
			From:        2,
			LastPage:    2,
			Path:        "",
			PerPage:     "10",
			To:          3,
			Total:       10,
			Links: []model.DynamicCommentMetaLink{
				model.DynamicCommentMetaLink{
					Url:    fmt.Sprintf("http://%s:8080/assets/post/post03.jpg", localAddr),
					Label:  "Beauty",
					Active: true,
				},
				model.DynamicCommentMetaLink{
					Url:    fmt.Sprintf("http://%s:8080/assets/post/post05.jpg", localAddr),
					Label:  "Beauty",
					Active: true,
				},
				model.DynamicCommentMetaLink{
					Url:    fmt.Sprintf("http://%s:8080/assets/post/post06.jpg", localAddr),
					Label:  "Beauty",
					Active: true,
				},
			},
		},
	}

	dynamicDetail = &model.DynamicDetailEntity{
		Data: model.DynamicDetailData{
			Id:           10000,
			UserId:       1,
			ThumbCount:   100,
			Thumbed:      55,
			Collected:    999,
			CollectCount: 1000,
			CommentCount: 1111,
			CreatedAt:    time.Now().Format("2006-0102 01:02:03"),
			UpdatedAt:    time.Now().Format("2006-0102 01:02:03"),
			Content:      "呵呵，天气这么好....",
			Location:     "Wuhan,Hubei,CN",
			User: model.DynamicDetailDataUser{
				Id:                1,
				Username:          "老痰涮菜",
				Nickname:          "老痰涮菜",
				Avatar:            fmt.Sprintf("http://%s:8080/assets/header03.png", localAddr),
				Email:             "gogogo@qq.com",
				Introduction:      "这个人很懒，什么也没留下...",
				NotificationCount: 10,
				Status:            1,
				CreatedAt:         "2009-09-09 09:09:09",
				UpdatedAt:         "2020-08-11 07:32:33",
			},
			Images: []model.DynamicDetailDataImage{
				model.DynamicDetailDataImage{
					Id:        12,
					UserId:    1,
					DynamicId: 10001,
					Type:      "image",
					Path:      fmt.Sprintf("http://%s:8080/assets/post/post08.jpg", localAddr),
				},
				model.DynamicDetailDataImage{
					Id:        12,
					UserId:    1,
					DynamicId: 10001,
					Type:      "image",
					Path:      fmt.Sprintf("http://%s:8080/assets/post/post07.jpg", localAddr),
				},
			},
		},
	}
	messageList = &model.MessageListEntity{
		Data: []model.MessageListData{
			model.MessageListData{
				Stranger: model.MessageListDataStranger{
					Id:                1,
					Username:          "lsp",
					Nickname:          "lyb",
					Online:            1,
					Avatar:            fmt.Sprintf("http://%s:8080/assets/header0%d.png", localAddr, rand.Int31()%9+1),
					CreatedAt:         time.Now().Format("2006-09-08 11:12:13"),
					UpdatedAt:         time.Now().Format("2006-09-08 11:12:13"),
					NotificationCount: 10,
					Status:            2,
				},
				Count:     1,
				Type:      "image",
				Id:        1,
				Excerpt:   "fgdfgdf",
				Draft:     "fresh meat",
				Createdat: time.Now().Format("2006-09-08 11:12:13"),
				Updatedat: time.Now().Format("2006-09-08 11:12:13"),
			},
		},
	}

}

func (m *MockCenter) GetMe() *model.User {

	return me
}

func (m *MockCenter) GetDynamicList() *model.DynamicListEntity {
	return dynamicList
}

func (m *MockCenter) GetDComments() *model.DynamicCommentEntity {
	return dcomments
}

func (m *MockCenter) GetDynamicDetail() *model.DynamicDetailEntity {
	return dynamicDetail
}

func (m *MockCenter) GetMessageList() *model.MessageListEntity {
	return messageList
}
