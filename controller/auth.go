package ctrl

import (
	"coward-saviour/mock"
	"coward-saviour/util"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func SendLoginCodes(c *gin.Context) {
	// phone := c.PostForm("phone")
	// log.Println(c.PostForm("phone"))
	key := fmt.Sprintf("verificationCode_%s", util.RandomString(15))

	// service.RedisCache.Set(key, [string]string{
	// 	key: {
	// 		"phone": phone,
	// 		"code":  "1234",
	// 	},
	// }, time.Minute*5)

	result := make(map[string]interface{})
	result["key"] = key
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

func Logout(c *gin.Context) {

}

func LoginByCode(c *gin.Context) {
	c.Request.ParseForm()
	verification_code, _ := c.GetPostForm("verification_code")
	verification_key, _ := c.GetPostForm("verification_key")
	log.Println("verify info:", verification_code, verification_key)

	// sendSms(code)
	// key := util.RandomString(15)
	expire_at := time.Now().Add(time.Minute * 5)
	result := make(map[string]interface{})
	result["data"] = mock.Mock.GetMe()
	result["access_token"] = expire_at.String()
	result["token_type"] = "Bearer"
	result["expired_in"] = time.Now().Add(time.Hour * 24).Format("2006-02-02 11:22:33")

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
