package main

import (
	"log"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("int", "./logagent.conf")
	if err != nil {
		log.Println("new config failed, err:", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		log.Println("read server:port failed, err:", err)
		return
	}

	log.Println("port:", port)
	logLevel := conf.String("logs::log_level")
	log.Println("log_level:", logLevel)

	logPath := conf.String("logs::log_path")
	log.Println("log_path:", logPath)
}
