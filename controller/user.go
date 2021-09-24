package ctrl

import (
	"coward-saviour/service"
	"coward-saviour/util"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbClient *mongo.Client

func UserInfo(c *gin.Context) {
	c.Request.ParseForm()

	util.RespOk(c.Writer, nil, "")
}

func LoginByCodes(c *gin.Context) {
	c.Request.ParseForm()
	verify_key := c.Request.Header.Get("verification_key")
	data, err := service.Cache.Get(verify_key)
	if err != nil || data == nil {
		log.Println(err)
		// c.Errors = append(c.Errors, )
	}
	var verify_data map[string]interface{}
	err = json.Unmarshal(data, verify_data)
	if err != nil {
		log.Println(err)
		// c.Errors = append(c.Errors, )
	}
	phone := verify_data["phone"]
	code := verify_data["code"]

	util.RespOk(c.Writer, nil, "")
}

func SendLoginCodes(c *gin.Context) {
	c.Request.ParseForm()
	phone := c.Params.ByName("phone")
	code := util.RandomString(4)
	sendSms(code)
	key := util.RandomString(15)
	expire_at := time.Now().Add(time.Minute * 5)
	result := make(map[string]interface{})
	result["key"] = key
	result["expire_at"] = expire_at.String()
	data, _ := json.Marshal(result)
	c.Writer.Write(data)
	// todo save {key:{phone,code},expireTime} to cache server
	var sdata map[string]interface{}
	sdata = make(map[string]interface{})
	sdata["phone"] = phone
	sdata["code"] = code
	data, _ = json.Marshal(sdata)
	service.Cache.Set(key, []byte(data))
}

func sendSms(code string) {
	//todo
}
