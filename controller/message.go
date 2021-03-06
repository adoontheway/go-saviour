package ctrl

import (
	"coward-saviour/mock"
	"coward-saviour/util"

	"github.com/gin-gonic/gin"
)

func InitMessage(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
func MessageList(c *gin.Context) {
	result := make(map[string]interface{})
	result["data"] = mock.Mock.GetMessageList()
	util.GinResp(c, 200, mock.Mock.GetMessageList(), "")
}
func RecordMessage(c *gin.Context) {
	//todo
	util.GinResp(c, 200, nil, "")
}
func SendMessage(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
func MessageStatus(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
func ReadMessage(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
func MessageWindow(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
func VideoMessage(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
func VideoMessageCallback(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
func MessageReadCallback(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
