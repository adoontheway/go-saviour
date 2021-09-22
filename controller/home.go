package ctrl

import (
	"coward-saviour/util"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
