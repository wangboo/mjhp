package mjhp

import (
	"github.com/astaxie/beego/config"
	"log"
)

type setting struct {
	IsDebugOn bool
	ComputeNum int
	KafkaTopic string
	KafkaAddr []string
}

var cfg *setting

func LoadConfig() {
	c, err := config.NewConfig("ini", "mjhp.conf")
	if err != nil {
		log.Fatalln("load file mjhp.conf failed: ", err)
	}
	cfg = &setting{}
	cfg.IsDebugOn = c.DefaultBool("debug", true)
	cfg.ComputeNum = c.DefaultInt("compute.num", 4)
	cfg.KafkaTopic = c.DefaultString("kafka.topic", "mjhp")
	cfg.KafkaAddr = c.DefaultStrings("kafka.addr", []string{"127.0.0.1:9092"})
}
