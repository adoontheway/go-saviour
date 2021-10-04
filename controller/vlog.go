package ctrl

import (
	"coward-saviour/util"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
func UploadFiles(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}

func VLogs(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func VLogCreate(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func VLogSave(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func VLogEdit(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func VLogUpdate(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
func VLogShow(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}
