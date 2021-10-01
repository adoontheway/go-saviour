package mock

import (
	"coward-saviour/model"
	"math/rand"
	"time"
)

var (
	me            *model.DUser
	dynamicList   *model.DynamicListEntity
	dcomments     *model.DynamicCommentEntity
	dynamicDetail *model.DynamicDetailEntity
)

type MockCenter struct {
}

func (m *MockCenter) Init() {
	me = &model.DUser{
		Id:            1,
		Username:      "test001",
		Nickname:      "test001",
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
		Email:         "123@163.com",
		Phone:         "13500000000",
		Status:        0,
		Introduction:  "我是一条酸菜鱼，又酸又菜又多余...",
		RememberToken: "12345678",
		Avatar:        "http://localhost:8080/assets/header01.png",
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
					Avatar:            "http://localhost:8080/assets/header01.png",
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
						Path:      "http://localhost:8080/assets/header01.png",
					},
				},
			},
		},
		Links: model.DynamicListLink{
			First: "http://localhost:8080/assets/header01.png",
			Last:  "http://localhost:8080/assets/header01.png",
			Next:  "http://localhost:8080/assets/header01.png",
		},
		Meta: model.DynamicListMeta{
			CurrentPage: 1,
			From:        1,
			LastPage:    1,
			Path:        "http://localhost:8080/assets/header01.png",
			PerPage:     "123",
			To:          2,
			Total:       10,
			Links: []model.DynamicListMetaLink{
				model.DynamicListMetaLink{
					Url:    "http://localhost:8080/assets/header01.png",
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
				ReplierAvatar:    "http://localhost:8080/assets/header02.png",
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
				ReplierAvatar:    "http://localhost:8080/assets/header04.png",
				ReceiverId:       4,
				ReceiverNickname: "nidaye",
				ThumbCount:       rand.Uint32() % 9999,
				Content:          "呵呵",
				CreatedAt:        time.Now().Format("2006-01-02 19:50:00"),
				UpdatedAt:        time.Now().Format("2013-01-02 19:50:00"),
			},
		},
		Links: model.DynamicCommentLink{
			First: "http://localhost:8080/assets/post/post01.jpg",
			Last:  "http://localhost:8080/assets/post/post04.jpg",
			Prev:  "http://localhost:8080/assets/post/post02.jpg",
			Next:  "http://localhost:8080/assets/post/post03.jpg",
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
					Url:    "http://localhost:8080/assets/post/post03.jpg",
					Label:  "Beauty",
					Active: true,
				},
				model.DynamicCommentMetaLink{
					Url:    "http://localhost:8080/assets/post/post05.jpg",
					Label:  "Beauty",
					Active: true,
				},
				model.DynamicCommentMetaLink{
					Url:    "http://localhost:8080/assets/post/post06.jpg",
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
				Avatar:            "http://localhost:8080/assets/header03.png",
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
					Path:      "http://localhost:8080/assets/post/post08.jpg",
				},
				model.DynamicDetailDataImage{
					Id:        12,
					UserId:    1,
					DynamicId: 10001,
					Type:      "image",
					Path:      "http://localhost:8080/assets/post/post07.jpg",
				},
			},
		},
	}

}

func (m *MockCenter) GetMe() *model.DUser {

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
