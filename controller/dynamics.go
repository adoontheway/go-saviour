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

	util.GinResp(c, 200, nil, "")
}
func DynamicCollect(c *gin.Context) {
	c.Request.ParseForm()
	util.GinResp(c, 200, nil, "")
}

func DynamicPublish(c *gin.Context) {
	c.Request.ParseForm()

	util.GinResp(c, 200, nil, "")
}
func DynamicInfo(c *gin.Context) {
	c.Request.ParseForm()

	util.GinResp(c, 200, mock.Mock.GetDynamicDetail(), "")
}
