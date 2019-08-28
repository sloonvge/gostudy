package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type res struct {
	r *http.Response
	err error
}


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan res, 1)
	req , err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		fmt.Println("http new request failed, err:", err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		pyload := res{r: resp, err:err}
		c <- pyload
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<- c
		fmt.Println("timeout")
	case r := <- c:
		defer r.r.Body.Close()
		out, _ := ioutil.ReadAll(r.r.Body)
		fmt.Println(out)
	}
}
