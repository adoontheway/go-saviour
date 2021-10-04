package ctrl

import (
	"coward-saviour/mock"
	"coward-saviour/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbClient *mongo.Client

func UserInfo(c *gin.Context) {
	c.Request.ParseForm()
	result := make(map[string]interface{})
	result["data"] = mock.Mock.GetMe()
	util.GinResp(c, 200, result, "")
}

func UploadAvatar(c *gin.Context) {
	util.GinResp(c, 200, nil, "")
}
