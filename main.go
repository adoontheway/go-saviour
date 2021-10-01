package main

import (
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
	r.GET("/users", ctrl.UserInfo)
	r.POST("/loginCodes", ctrl.SendLoginCodes)
	r.POST("/loginByCode", ctrl.LoginByCode)

	r.GET("/dynamics", ctrl.Dynamic)
	r.GET("/dynamics/create", ctrl.DynamicCreate)
	r.POST("/dynamics", ctrl.DynamicInfo)
	r.GET("/dynamics/{id}/edit", ctrl.DynamicEdit)
	r.PUT("/dynamics/{id}/", ctrl.DynamicUpdate)
	r.GET("/dynamics/{id}", ctrl.DynamicShow)

	r.POST("/dynamics/thumb", ctrl.DynamicThumb)
	r.POST("/dynamics/collect", ctrl.DynamicCollect)
	r.GET("/dcomments", ctrl.DComments)

	r.GET("/comments", ctrl.Comments)
	r.GET("/comments/create", ctrl.CommentCreate)
	r.POST("/comments", ctrl.CommentSave)
	r.GET("/comments/{id}/edit", ctrl.CommentEdit)
	r.PUT("/comments/{id}/", ctrl.CommentUpdate)
	r.GET("/comments/{id}", ctrl.CommentShow)

	r.GET("/vlogs", ctrl.VLogs)
	r.GET("/vlogs/create", ctrl.VLogCreate)
	r.POST("/vlogs", ctrl.VLogSave)
	r.GET("/vlogs/{id}/edit", ctrl.VLogEdit)
	r.PUT("/vlogs/{id}/", ctrl.VLogUpdate)
	r.GET("/vlogs/{id}", ctrl.VLogShow)
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
