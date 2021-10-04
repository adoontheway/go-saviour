package main

import (
	constants "coward-saviour/const"
	ctrl "coward-saviour/controller"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//cors middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", " OPTIONS"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"content-length"},
		AllowOriginFunc: func(origin string) bool {
			log.Println("request from origin:", origin)
			return true
		},
	}))
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	r.GET("/", ctrl.Index)

	r.POST(constants.LoginCodes, ctrl.SendLoginCodes)
	r.POST(constants.LoginByCode, ctrl.LoginByCode)
	r.POST(constants.UserLogout, ctrl.Logout)
	r.GET(constants.UserInfo, ctrl.UserInfo)

	r.POST(constants.UploadAvatar, ctrl.UploadAvatar)
	r.POST(constants.UploadFile, ctrl.UploadFiles)
	r.POST(constants.UploadImage, ctrl.UploadImg)
	r.GET(constants.DynamicComment, ctrl.DComments)

	r.GET(constants.DynamicList, ctrl.Dynamic)
	r.POST(constants.ThumbDynamic, ctrl.DynamicThumb)
	r.POST(constants.CollectDynamic, ctrl.DynamicCollect)
	r.GET(constants.PublishDynamic, ctrl.DynamicPublish)
	r.POST(constants.DynamicList, ctrl.DynamicInfo)

	r.GET(constants.Comment, ctrl.Comments)
	r.POST(constants.Comment, ctrl.PubComment)

	r.POST(constants.InitMessage, ctrl.InitMessage)
	r.POST(constants.MessageList, ctrl.MessageList)
	r.POST(constants.SendMessage, ctrl.SendMessage)
	r.GET(constants.RecordMessage, ctrl.RecordMessage)
	r.POST(constants.MessageStatus, ctrl.MessageStatus)
	r.POST(constants.ReadMessage, ctrl.ReadMessage)
	r.POST(constants.MessageWindow, ctrl.MessageWindow)
	r.POST(constants.VideoMessage, ctrl.VideoMessage)
	r.POST(constants.VideoMessageCallback, ctrl.VideoMessageCallback)
	r.POST(constants.MessageReadCallback, ctrl.MessageReadCallback)

	r.StaticFS("/assets", http.Dir("assets"))

	r.Run(":8080")

	// http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	// http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	// http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	// http.HandleFunc("/contact/createcommunity", ctrl.CreateCommunity)
	// http.HandleFunc("/contact/addfriend", ctrl.AddFriend)
	// http.HandleFunc("/chat", ctrl.Chat)
	// http.HandleFunc("/attach/upload", ctrl.Upload)

	// http.Handle("/asset/", http.FileServer(http.Dir(".")))
	// http.Handle("/mnt/", http.FileServer(http.Dir(".")))
	// http.ListenAndServe(":8080", nil)
}
