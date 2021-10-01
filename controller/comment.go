package ctrl

import (
	"coward-saviour/mock"
	"coward-saviour/util"

	"github.com/gin-gonic/gin"
)

func DComments(c *gin.Context) {
	c.Request.ParseForm()
	// result := make(map[string]interface{})
	// result["data"] = mock.Mock.GetDComments()
	util.GinResp(c, 200, mock.Mock.GetDComments(), "")
}

func Comments(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func CommentCreate(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func CommentSave(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func CommentEdit(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func CommentUpdate(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func CommentShow(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
