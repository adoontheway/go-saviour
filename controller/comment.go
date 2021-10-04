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

func PubComment(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
