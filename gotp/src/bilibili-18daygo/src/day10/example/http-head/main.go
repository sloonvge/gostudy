package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var urls = []string{
		"http://www.baidu.com",
		"http://google.com",
		"http://taobao.com",
	}

	client := &http.Client{
		// Transport: &http.Transport{
		// 	Dial: func(network, addr string) (conn net.Conn, e error) {
		// 		return net.DialTimeout(network, addr, time.Second)
		// 	},
		// },
		Timeout: time.Second,
	}
	for _, url := range urls{
		resp, err := client.Head(url)
		if err != nil {
			fmt.Printf("head %s err, %v\n", url, err)
			continue
		}

		fmt.Println(resp.Status)
	}
}