package ctrl

import (
	"coward-saviour/mock"
	"coward-saviour/util"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func SendLoginCodes(c *gin.Context) {
	log.Println(c.PostForm("phone"))
	result := make(map[string]interface{})
	result["key"] = "1234"
	result["expired_at"] = time.Now().Add(time.Minute * 5)
	util.GinResp(c, 200, result, "")
	// phone := c.Params.ByName("phone")
	// code := util.RandomString(4)
	// sendSms(code)
	// key := util.RandomString(15)
	// expire_at := time.Now().Add(time.Minute * 5)
	// result := make(map[string]interface{})
	// result["key"] = key
	// result["expire_at"] = expire_at.String()
	// data, _ := json.Marshal(result)
	// c.Writer.Write(data)
	// // todo save {key:{phone,code},expireTime} to cache server
	// var sdata map[string]interface{}
	// sdata = make(map[string]interface{})
	// sdata["phone"] = phone
	// sdata["code"] = code
	// data, _ = json.Marshal(sdata)
	// service.Cache.Set(key, []byte(data))
}

func LoginByCode(c *gin.Context) {

	c.Request.ParseForm()
	verification_code, _ := c.GetPostForm("verification_code")
	log.Println("verify code", verification_code)

	// sendSms(code)
	// key := util.RandomString(15)
	expire_at := time.Now().Add(time.Minute * 5)
	result := make(map[string]interface{})
	result["data"] = mock.Mock.GetMe()
	result["access_token"] = expire_at.String()
	result["token_type"] = "Bearer"
	result["expired_in"] = time.Now().Add(time.Hour * 24)

	util.GinResp(c, 200, result, "")
	// todo save {key:{phone,code},expireTime} to cache server
	// var sdata map[string]interface{}
	// sdata = make(map[string]interface{})
	// sdata["phone"] = phone
	// sdata["code"] = code
	// data, _ = json.Marshal(sdata)
	// service.Cache.Set(key, []byte(data))
}

func sendSms(code string) {
	//todo
}
