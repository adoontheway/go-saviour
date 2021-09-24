package setting

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var GlobalSetting *Setting

type Setting struct {
	Mongo MongoDb
}

type MongoDb struct {
	Host        string
	Username    string
	Password    string
	Timeout     int32
	Maxpoolsize int32
	Heartbeat   int32
}

func InitSetting() {
	file, err := ioutil.ReadFile("./conf/app.yml")
	if err != nil {
		log.Fatal("failed to read file:", err)
	}
	err = yaml.Unmarshal(file, GlobalSetting)
	if err != nil {
		log.Fatal("failed to parse yaml:", err)
	}
}
