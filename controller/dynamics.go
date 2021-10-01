package ctrl

import (
	"coward-saviour/mock"
	"coward-saviour/util"

	"github.com/gin-gonic/gin"
)

func Dynamic(c *gin.Context) {
	c.Request.ParseForm()
	dynamic := mock.Mock.GetDynamicList()
	util.GinResp(c, 200, dynamic, "")
}

func DynamicThumb(c *gin.Context) {
	c.Request.ParseForm()
	dynamic := mock.Mock.GetDynamicList()
	util.GinResp(c, 200, dynamic, "")
}
func DynamicCollect(c *gin.Context) {
	c.Request.ParseForm()
	dynamic := mock.Mock.GetDynamicList()
	util.GinResp(c, 200, dynamic, "")
}

func DynamicCreate(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func DynamicInfo(c *gin.Context) {
	c.Request.ParseForm()

	util.GinResp(c, 200, mock.Mock.GetDynamicDetail(), "")
}
func DynamicEdit(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func DynamicUpdate(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func DynamicShow(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
