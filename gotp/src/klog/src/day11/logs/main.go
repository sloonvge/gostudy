package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logs/test.log"
	config["level"] = logs.LevelInfo

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("debug, info1")
	logs.Trace("trace, info2")
	logs.Warn("warn, info3")
}