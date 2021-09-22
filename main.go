package main

import (
	ctrl "coward-saviour/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", ctrl.Index)
	r.GET("/users", ctrl.UserInfo)

	r.GET("/dynamic", ctrl.Dynamic)
	r.GET("/dynamic/create", ctrl.DynamicCreate)
	r.POST("/dynamic", ctrl.DynamicSave)
	r.GET("/dynamic/{id}/edit", ctrl.DynamicEdit)
	r.PUT("/dynamic/{id}/", ctrl.DynamicUpdate)
	r.GET("/dynamic/{id}", ctrl.DynamicShow)

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

	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/createcommunity", ctrl.CreateCommunity)
	http.HandleFunc("/contact/addfriend", ctrl.AddFriend)
	http.HandleFunc("/chat", ctrl.Chat)
	http.HandleFunc("/attach/upload", ctrl.Upload)

	// http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
